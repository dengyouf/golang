package basic

import (
	"fmt"
	"testing"
)

func TestPointerRaw(t *testing.T) {
	a := "string a"
	fmt.Println(&a) //0xc0000889c0
}

func TestPointer1(t *testing.T) {
	var a *int
	fmt.Println(a) // nil
}

func TestPointer11(t *testing.T) {
	var a *int
	a = new(int)
	*a = 100
	fmt.Println(a) // nil
}

func TestPPP(t *testing.T) {
	var a ****int

	v := 10
	p1 := &v
	p2 := &p1
	p3 := &p2
	a = &p3
	fmt.Println(v, p1, p2, p3, a)
}

func TestPointer12(t *testing.T) {
	var ip *int = new(int)
	fmt.Println(&ip, ip, *ip)
}
