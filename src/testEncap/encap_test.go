package testEncap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestEncap(t *testing.T) {
	e := Employee{"1", "kaikai", 25}
	e1 := Employee{Name:"kaikai", Age:20}
	e2 := new(Employee)
	e2.Age = 25
	e2.Name = "kaikaikai"
	e2.Id = "2"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}

func (e Employee) Get() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	e.Name = "peipei"
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) Get2() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	e.Name = "peipei"
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestObjFunc(t *testing.T) {
	e := Employee{"1", "kaikai", 25}
	t.Log(e)
	t.Log(e.Get())
	t.Log(e)
}