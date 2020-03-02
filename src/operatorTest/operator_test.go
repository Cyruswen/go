package operatorTest

import "testing"

func TestArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 2} //顺序不同, 可以比较, 结果是false
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	//e := [...]string{"a", "b", "c"}
	//t.Log(a == e) //不同类型, 不可以比较
	t.Log(a == b)
	//t.Log(a == c) golang长度不同的数组, 不能比较
	//只要长度相同, 类型相同的数组就可以比较
	t.Log(a == d)
}

func TestClear(t *testing.T) {
	a := 15
	b := 1
	c := 0
	d := 7
	t.Log(a &^ 4, a &^ 2)   // &^ 0 保持不变
	t.Log(b &^ 1, b &^ 0)
	t.Log(c &^ 1, c &^ 0)
	t.Log(d &^ 2, d &^ 1)
}
