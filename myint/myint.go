package myint

// MyInt is a ....
type MyInt int32

// Add is adding from the given param an return a copy.
func (i MyInt) Add(elem int) MyInt {
	return i + MyInt(elem)
}

func (i MyInt) Sub(elem int) MyInt {
	return i - MyInt(elem)
}

func (i MyInt) Multiply(elem int) MyInt {
	return i * MyInt(elem)
}

func (i MyInt) Divide(elem int) MyInt {
	return i / MyInt(elem)
}
