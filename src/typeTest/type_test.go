package typeTest

import (
	"math"
	"testing"
)

type myInt int64

func TestType(t *testing.T) {
	//var a string = "hello world";
	var a int64 = 1
	var b int
	//b = int(a)
	var c myInt = 2
	b = int(c)
	t.Log(a, b, c)
}

func TestMaxType(t *testing.T) {
	var a float32 = math.MaxFloat32;
	//b := math.MaxFloat64
	aPtr := &a
	//aPtr = aPtr + 1  Go不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T, %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if (s == "") {
		t.Log(s + "=\"\"")
	}
	//Golang没有++a
}
