package basic

import (
	"fmt"
	"testing"
)

var (
	name string
	age  uint
)

func TestSscan(t *testing.T) {
	input := "dengyou 30"
	fmt.Sscan(input, &name, &age)
	fmt.Println(name, age)
}

func TestSscan2(t *testing.T) {
	input := "dengyou 30 40"
	fmt.Sscan(input, &name, &age)
	fmt.Println(name, age) //dengyou 30
}

func TestSscan3(t *testing.T) {
	input := "dengyou\n30"
	fmt.Sscan(input, &name, &age)
	fmt.Println(name, age) //dengyou 30
}

func TestSscanln(t *testing.T) {
	input := "dengyou 30 40"
	fmt.Sscanln(input, &name, &age)
	fmt.Println(name, age) //dengyou 30
}

func TestSscanln2(t *testing.T) {
	input := "dengyou \n 30 40"
	fmt.Sscanln(input, &name, &age)
	fmt.Println(name, age) //dengyou 30
}

func TestSscanf(t *testing.T) {
	input := "dengyou : 50"
	fmt.Sscanf(input, "%s : %d", &name, &age)
	fmt.Println(name, age) //dengyou 50
}

func TestSscanRet(t *testing.T) {
	input := "dengyou : 50"
	ret, err := fmt.Sscanf(input, "%s : %d", &name, &age)
	fmt.Println(ret, err)
}

func TestSscanRet2(t *testing.T) {
	input := "andy : err"
	ret, err := fmt.Sscanf(input, "%s : %d", &name, &age)
	fmt.Println(ret, err)
}
