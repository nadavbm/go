package math

import (
	"testing"
)

type TestSome struct {
	nums   []int
	result int
}

func TestGetSum(t *testing.T) {
	xy := GetSum(1, 2, 3)
	if xy != 6 {
		t.Error("Should be ", 6, "But its", xy)
	}

	yz := []TestSome{
		TestSome{[]int{10, 20, 30}, 60},
	}

	for _, v := range yz {
		zz := GetSum(v.nums...)
		if zz != v.result {
			t.Error("expected", v.result, "got", zz)
		}
	}
}
