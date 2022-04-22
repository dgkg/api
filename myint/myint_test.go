package myint_test

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"

	"github.com/dgkg/api/myint"
)

func TestMyIntAdd(t *testing.T) {
	data := []struct {
		title           string
		value           myint.MyInt
		param           int
		should          myint.MyInt
		isWaitingForErr bool
		err             error
	}{
		{"A", 1, 1, 2, false, nil},
		{"B", math.MaxInt32, 1, 0, true, myint.ErrOutOfRange},
		{"C", 9, 1, 10, false, nil},
	}
	for _, v := range data {
		mi := v.value
		got, _ := mi.Add(v.param)
		if got != v.should {
			t.Error("for", v.title, "got", got, "should got", v.should)
		}
	}
}

func TestMyIntSub(t *testing.T) {
	assert := assert.New(t)
	var a myint.MyInt = 10
	var b int = 5
	var c myint.MyInt = 5

	assert.Equal(a.Sub(b), c, "The result should be 5.")
}

func TestMyIntMultiply(t *testing.T) {
	Convey("2 multiply by 2 should equal 4", t, func() {
		So(myint.MyInt(2).Multiply(2), ShouldEqual, 4)
	})
}

func TestMyIntDivide(t *testing.T) {
	data := []struct {
		title           string
		value           myint.MyInt
		param           int
		should          myint.MyInt
		isWaitingForErr bool
		err             error
	}{
		{title: "A", value: 1, param: 1, should: 1, isWaitingForErr: false, err: nil},
		{"try to divide by zero", 9, 0, 0, true, myint.ErrDivideByZero},
	}
	for _, v := range data {
		mi := v.value
		got, err := mi.Divide(v.param)
		if err == nil && v.isWaitingForErr {
			t.Error("for", v.title, "got no error but should have :", v.err)
		}

		if err != v.err {
			t.Error("for", v.title, "waiting for err", v.err, "but got", err)
		}

		if got != v.should {
			t.Error("for", v.title, "got", got, "should got", v.should)
		}
	}
}

func TestMyIntDivide2(t *testing.T) {
	assert := assert.New(t)
	var a myint.MyInt = 10
	var b int = 0
	var waitingFor myint.MyInt = 0

	res, err := a.Divide(b)
	assert.Equal(res, waitingFor, "The result should be 0.")
	assert.Equal(err, myint.ErrDivideByZero, "The result should be an error.")
}
