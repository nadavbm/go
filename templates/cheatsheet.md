### templates

templates can pass data from go app to text or html templates

the template file can be called in any name suffix. usually named ```tpl.gohtml```

#### text templates

to import text templates use ```import "text/template"```

to use a specific text template in go -

1. parse the template:

```
tpl, err := template.ParseFiles("tpl.gohtml")
if err != nil {
    log.Fatalln(err)
}
```

2. execute the template:
```
err = tpl.Execute(os.Stdout, nil)
if err != nil {
    log.Fatalln(err)
}
```

to use a templates folder -

1. create a directory `cd project/ && mkdir templates`

2. add the template files `touch template/tpl1.gotmpl`

3. use `template.ParseGlob("templates/*")`

4. excute the templates

```
tpl, err := template.ParseGlob("templates/*")
if err != nil {
    log.Fatalln(err)
}

err = tpl.Excute(os.Stdout, nil)
if err != nil {
    log.Fatalln(err)
}

err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gotmpl", nil)
if err != nil {
    log.Fatalln(err)
}
```

we can also initialize glob of templates with init func and use `template.Must` to parse this folder only once:

```
var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("templates/*"))
}
```

##### passing data to a tempalte

to pass data use this syntax inside the template ```{{.}}```. the `.` is the piece of data which we pass in.

for example passing 12 (an int) - in template `{{.}}` will be `12`:

```
err = tpl.ExecuteTemplate(os.Stdout, "tpl1.gotmpl", 12)
if err != nil {
    log.Fatalln(err)
}
```

```{{$varName := .}}``` - assigning a value to a variable in a template

to use this var - ```{{$varName}}```

###### composite data

passing a slice ```[]slice{"bunch","0f","strings"}``` to a template and ranging this slice in the template:

```
<ul>
    {{range .}}
    <li>{{.}}</li>
    {{end}}
</ul>
```

using index and elements as a variables:

```
<ul>
    {{range $i, $e := .}}
    <li>{{index: $i, element: $e}}</li>
    {{end}}
</ul>
```

passing a map to a template:

```
album := map[string]string{
    "Mastodon": "Leviathan",
    "Faith No More": "Album of the year",
}
```
and using key value in the template:
```
<ul>
    {{range $k, $v := .}}
    <li>{{$k : $v}}</li>
    {{end}}
</ul>
```

passing struct fields to a template:

```
type Person struct{
    Name    string
    Age     int
}

func main() {
    tavori := Person{
        Name: "shimi",
        Age: 13,
    }

    err := tpl.Execute(os.Stdout, tavori)
    if err != nil {
        log.Fatalln(err)
    }
}
```

in tpl1.gotmpl:

```
<h1>Hello {{.Name}}</h1>
<p>Your age is {{.Age}}</p>
```

##### pass functions to templates

create functions in your go code and assign it to a template:
```
var tpl *template.Template

var funcMap = tempalte.FuncMap{
    "getSum": getSum,
}

func init() {
    tpl = template.Must(tempalte.New("").Funcs(getSum).ParseFiles("tpl3.gotmpl"))
}

func getSum(xy ...int) int {
	sum := 0
	for _, i := range xy {
		sum += i
	}
	return sum
}

func main() {
    numbers := []int{1,23,42,1115,1324312}
    err := tpl.Execute(os.Stdout, numbers)
    if err != nil {
        log.Fatalln(err)
    }
}
```
in tpl3.gotmpl:
```
<h1>Sum calculator</h1>
<p>Your sum is {{getSum .}}</p>
```


#### html templates