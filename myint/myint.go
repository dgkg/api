package myint

import (
	"errors"
	"math"
)

// MyInt is a ....
type MyInt int32

var (
	ErrDivideByZero = errors.New("myint: try to divide by zero")
	ErrOutOfRange   = errors.New("myint: out of range")
)

// Add is adding from the given param an return a copy.
func (i MyInt) Add(elem int) (MyInt, error) {
	if elem > math.MaxInt32 {
		return 0, ErrOutOfRange
	}
	maxValueToAdd := math.MaxInt32 - i
	if maxValueToAdd > MyInt(elem) {
		return i + MyInt(elem), nil
	}
	return 0, ErrOutOfRange
}

func (i MyInt) Sub(elem int) MyInt {
	return i - MyInt(elem)
}

func (i MyInt) Multiply(elem int) MyInt {
	return i * MyInt(elem)
}

func (i MyInt) Divide(elem int) (MyInt, error) {
	if elem == 0 {
		return 0, ErrDivideByZero
	}
	return i / MyInt(elem), nil
}
