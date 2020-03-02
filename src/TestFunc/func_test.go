package TestFunc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//1. 可以有多个返回值
//2. 所有参数都是值传递: slice, map, channel也是
//3. 函数可以作为变量的值
//4. 函数可以作为参数和返回值

func TestFunc(t *testing.T){
	a, b := returnMultiValues()
	t.Log(a, b)
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20);
}

func TestFn(t *testing.T) {
	tssf := timeSpent(slowFunc)
	ret := tssf(10)
	t.Log(ret)
}

func slowFunc(op int) int{
	time.Sleep(time.Second * 1)
	return op
}

func timeSpent(inner func(op int) int) func (op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer(t *testing.T) {
	defer Clear() //释放资源, 释放锁, 延迟执行
	fmt.Println("Start")
	panic("err")
}
