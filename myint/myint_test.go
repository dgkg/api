package myint_test

import (
	"testing"

	"github.com/dgkg/api/myint"
)

func TestMyIntAdd(t *testing.T) {
	data := []struct {
		title  string
		value  myint.MyInt
		param  int
		should myint.MyInt
	}{
		{"A", 1, 1, 2},
		{"B", 2, 1, 3},
		{"C", 9, 1, 10},
	}
	for _, v := range data {
		mi := v.value
		got := mi.Add(v.param)
		if got != v.should {
			t.Error("for", v.title, "got", got, "should got", v.should)
		}
	}

}

func TestMyIntSub(t *testing.T) {

}

func TestMyIntMultiply(t *testing.T) {

}

func TestMyIntDivide(t *testing.T) {

}
