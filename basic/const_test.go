package basic

import (
	"fmt"
	"testing"
)

const (
	a = 10
	b = iota
	c
	d = 50
	e = iota
	f
)

func TestConstPrint(t *testing.T) {
	fmt.Println(a, b, c, d, e, f) // 10 1 2 50 4 5
}

func TestLocalConst(t *testing.T) {
	const a, b = 10, 20
	fmt.Println(a, b)
}
