package fib

import (
	"fmt"
	"testing"
)

/*
func TestFibList(t *testing.T){
	a := 1
	b := 1
	for i := 0; i < 5; i++ {
		t.Log(a)
		tmp := a
		a = b
		b = b + tmp
	}
}*/


func TestExchange(t *testing.T) {
	a := 1;
	b := 2;
	fmt.Println(a, " ", b)
	a, b = b, a
	fmt.Println(a, " ", b)
}

