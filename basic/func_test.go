package basic

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestAddFunc(t *testing.T) {
	x, y := 3, 4

	res := add(x, y)
	fmt.Println(res)

}
func add(x, y int) int {
	return x + y
}

func TestSwapPtr(t *testing.T) {
	x, y := 4, 5
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x, y *int) {
	*x, *y = *y, *x
	return
}

func max(a, b int, args ...int) int {
	var max_value int

	if a > b {
		max_value = a
	} else {
		max_value = b
	}

	for _, value := range args {
		if max_value < value {
			max_value = value
		}
	}
	return max_value
}

func TestMaxValue(t *testing.T) {

	fmt.Println(max(1, 3, 6, 2, 4, 7))
}

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func fact2(n int) int {
	current := n
	for n > 1 {
		fmt.Printf("%d * (%d-1)\n", current, n)
		current = current * (n - 1)
		n--
	}
	return current
}
func TestFact(t *testing.T) {
	fmt.Println(fact(5))
	fmt.Println(fact2(5))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func TestFib(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, fib(i))
	}
	fmt.Println(s)
}

func ParseInt(str string) (result int64, err error) {
	result = 10
	return 20, nil
}

func TestRetrun(t *testing.T) {
	fmt.Println(ParseInt(""))
}

func TestReturnHull(t *testing.T) {
	i, _ := strconv.ParseInt("10", 10, 64)
	fmt.Println(i)
}

func TestAnonymousFunc(t *testing.T) {
	a := func(x, y int) int {
		return x + y
	}

	fmt.Println(a(1, 2))
}

func TestAnonymousFunc2(t *testing.T) {
	a := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(a)
}

func TestFuncType(t *testing.T) {
	addFunc := func(x, y int) int {
		return x + y
	}
	fmt.Println(reflect.TypeOf(addFunc))
	fmt.Println(addFunc(10, 20))
}

type addFunc func(x, y int) int

func asArg(fn addFunc) int {
	return fn(2, 2) * 2
}

func TestFuncAsArg(t *testing.T) {
	ret := asArg(func(x, y int) int {
		return x + y
	})
	fmt.Println(ret)
}

func TestFun5(t *testing.T) {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Printf("%T\n", add) // func(int, int) int
	add(1, 3)

	m := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(m)
}

func TestTT(t *testing.T) {
	fn := TT()
	fmt.Println(fn())
}

func TT() func() int {
	x, y := 10, 20 // GC  a -> x,  b -> y  ref 1

	add := func() int {
		return x + y
	}

	return add // 10
}
