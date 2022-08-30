package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArithmetic(t *testing.T) {
	a := 10
	b := 12.5
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))

	// 不同类型是不可以计算的  (mismatched types int and float64)
	// res := a + b
	// fmt.Println(res)

	// 强制类型转换
	res := float64(a) + b
	fmt.Println(res)

	t.Log(a + int(b))
}

func TestSpecOp(t *testing.T) {
	a := 10
	a++
	fmt.Println(a)
	a--
	fmt.Println(a)
	a--
	fmt.Println(a)
	a--
	fmt.Println(a)
	t.Log(a) // /Users/dengyou/Desktop/golang/basic/op_test.go:35: 8
}

func TestBoolExpr(t *testing.T) {
	a, b := 10, 20
	fmt.Println(a > b)
	fmt.Println(a < b)

	if a > b {
		a, b = b, a
	}
	fmt.Println(a, b)
}

func TestBitOp(t *testing.T) {
	var (
		a uint = 1024
		b uint = 85
	)
	fmt.Printf("a&b = %d\n", a&b)
	fmt.Printf("a|b = %d\n", a|b)

	fmt.Printf("a>>2 = %d\n", a>>2)
	fmt.Printf("a<<2 = %d\n", a<<2)
}

func TestPointer(t *testing.T) {
	a := 10
	b := &a
	*b += 5
	fmt.Println(a)
}
