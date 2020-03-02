package TestMap

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"a":1, "b":2, "c":3}
	t.Logf("len m=%d", len(m))
	m1 := map[int]int{}
	m1[0] = 10
	t.Log(m1)
	m2 := make(map[int]int, 10)
	t.Logf("len m2=%d", len(m2))
}

func TestExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2]) //未定义的值是0 ,正常赋值的也是0值, 这种情况需要区分
	m1[3] = 0
	if v,ok := m1[3];ok {
		t.Log("key 3 is exist", v, ok)
	} else {
		t.Log("key 3 is not exist", v, ok)
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"a":1, "b":2}
	for k,v := range m1 {
		t.Log(k, v)
	}
}

func TestMapFunc(t *testing.T) {
	m := map[int]func(op int)int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapAsSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if n := 3; mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	if n := 1; mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}