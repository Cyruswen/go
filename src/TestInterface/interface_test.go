package TestInterface

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

type GoPragrammer struct {

}

func (pr *GoPragrammer) WriteHelloWorld() string {
	return  "fmt.Println(\"hello world\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoPragrammer)
	t.Log(p.WriteHelloWorld())
}