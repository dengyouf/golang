package basic

import (
	"fmt"
	"testing"
)

func TestMapDel(t *testing.T) {
	var m map[string]int
	fmt.Printf("%p\n", m)
	fmt.Println(m)

	m = make(map[string]int)
	fmt.Println(m)
	m = map[string]int{
		"tom":   1,
		"jerry": 2,
	}
	fmt.Println(m)
}

func TestMapDef(t *testing.T) {
	m := map[string]int{"Tony": 22, "Andy": 55}
	fmt.Println(m)
}

func TestMapDef2(t *testing.T) {
	a := make(map[int]string, 1)
	a[0] = "a"
	a[1] = "b"
	fmt.Println(a)
	fmt.Println(len(a))
}

type AddFunc func(x, y int) int

func TestMapDel3(t *testing.T) {
	var m1 map[string]string
	var m2 map[int]int
	var m3 map[float64]string
	var m4 map[string]AddFunc
	fmt.Println(m1, m2, m3, m4)
}

func TestMapDel4(t *testing.T) {
	var a map[int]string
	var b []string
	fmt.Printf("%p, %p\n", a, b)
	// a[0] = "a"
	// b[0] = "a"

	a = map[int]string{0: "a"}
	b = []string{"a"}
	fmt.Printf("%p, %p\n", a, b)
}

func TestMapAcc1(t *testing.T) {
	a := map[int]string{}
	fmt.Println(a[0])

	a[100] = "t100"
	v1, ok1 := a[100]
	fmt.Println(v1, ok1) //t100 true

	v2, ok2 := a[0]
	fmt.Printf("%q, %v", v2, ok2) //"", false
}

func TestMapAcc2(t *testing.T) {
	a := map[int]string{}
	fmt.Println(a[0])

	a[100] = "t100"

	if v, ok := a[100]; ok {
		fmt.Println(v)
	}

	fmt.Println(a)

	v, ok := a[1]
	if ok {
		fmt.Println(v)
	} else {
		a[1] = "t1"
	}
	fmt.Println(a)
}

func TestMapDel1(t *testing.T) {
	a := map[int]string{100: "t100", 1: "t1"}
	delete(a, 1)
	fmt.Println(a)
}

func TestMapAsFuncArg(t *testing.T) {
	a := map[int]string{1: "a", 2: "b", 3: "c"}

	func(map[int]string) {
		delete(a, 1)
	}(a)

	fmt.Println(a)
}

type Add func(x, y int) int

func TestMapFuncV(t *testing.T) {
	op := map[string]func(x, y int) int{
		"+": func(x, y int) int {
			return x + y
		},
		"-": func(x, y int) int {
			return x - y
		},
		"*": func(x, y int) int {
			return x * y
		},
		"/": func(x, y int) int {
			return x / y
		},
	}

	fmt.Println(op["+"](1, 2))
	fmt.Println(op["-"](1, 2))
}
