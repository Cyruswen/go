package constant

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1<<iota
	Writable
	Executable
)

func TestConstant(t *testing.T) {
	t.Log(Wednesday, Tuesday)
}

func TestConstant1(t *testing.T) {
	a := 7 //0111
	t.Log(a&Readable, a&Writable, a&Executable)
	t.Log(Readable, Writable, Executable)
}