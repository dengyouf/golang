package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVar(t *testing.T) {
	var a int
	var b int

	fmt.Println(a, b)
	a = 10
	b = 10
	fmt.Println(a, b)
}

func TestVarOnline(t *testing.T) {
	var a, b int = 8, 9
	fmt.Println(a, b)
}

func TestVarMutiline(t *testing.T) {
	var (
		a int = 9
		b int = 10
	)
	fmt.Println(a, b)
}

func TestVarShort(t *testing.T) {
	a := "string"
	b := true
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), a, b)
}

func TestVarQuota(t *testing.T) {
	a := 'a'
	b := "b"
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	fmt.Printf("%c, %s", a, b)
}
