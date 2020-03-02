package testExtend

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	//d.Speak()
	fmt.Println("wang")
}
//
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
}


type code string

type Pro interface {
	writeHello() code
}

type Gpro struct {
}

func (g *Gpro)writeHello() code {
	return "fmt.Printlen(\"hello\")"
}

type Jpro struct {
}

func (j *Jpro)writeHello() code {
	return "System.out.Printlen(\"hello\")"
}

type Hpro struct {
}

func (h *Hpro)writeHello() code {
	return "echo 'hello'"
}

func Poly(p Pro)  {
	fmt.Println(p.writeHello())
}

func TestPoly(t *testing.T) {
	j := &Jpro{}
	h := new(Hpro)
	g := new(Gpro)
	Poly(j)
	Poly(h)
	Poly(g)
}