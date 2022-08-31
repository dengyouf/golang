package basic

import (
	"fmt"
	"testing"
)

func TestForBase(t *testing.T) {
	var sum int = 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func TestForShort(t *testing.T) {
	var sum int
	for sum < 10 {
		sum++
	}
	fmt.Println(sum)
}

func TestForLoop(t *testing.T) {
	var sum int
	for {
		sum++
		fmt.Println(sum)

		if sum == 10 {
			return
		}
	}
}

func TestForRange1(t *testing.T) {
	iter := "abcde"
	for k, v := range iter {
		fmt.Printf("%d %c\n", k, v)
	}

	iter2 := "我爱北京天安门，yes"
	for i, v := range iter2 {
		fmt.Printf("%d %c\n", i, v)
	}
}

func TestForRange2(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range nums {
		fmt.Println(i, v)
		nums[i] = 99
	}
	fmt.Println(nums)
}

func TestFor99(t *testing.T) {
	for m := 1; m < 10; m++ {
		for n := 1; n < 10; n++ {
			if n <= m {
				fmt.Printf("%d * %d = %d\t", n, m, n*m)
			}
		}
		fmt.Println()
	}
}

func TestForPromise(t *testing.T) {
	for m := 2; m < 100; m++ {
		isP := true
		for n := 2; n <= (m / n); n++ {
			if m%n == 0 {
				isP = false
				break
			}
		}
		if isP {
			fmt.Println(m)
		}
	}

}

func TestBreak(t *testing.T) {
	i := 0
	for {
		fmt.Print(i, " ")
		if i == 5 {
			break
		}
		i++
	}
}

func TestBreakLabel(t *testing.T) {

	fmt.Println("--------break-------")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
	}
	for i2 := 11; i2 <= 13; i2++ {
		fmt.Printf("i2: %d\n", i2)
		break
	}

	fmt.Println("------break labels ------")

re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)

		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}
}

func TestContinue(t *testing.T) {
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

func TestLabelNext(t *testing.T) {
	i := 0
	for {
		if i > 5 {
			goto LABEL
		}
		fmt.Printf("%d :--------------\n", i)
		i++
	}

LABEL:
	fmt.Println("I will Finished~")
	return
}

func TestLabelLOOP(t *testing.T) {
	var a int = 10

LOOP:
	for a < 20 {
		if a == 15 {
			a++
			goto LOOP
		}
		fmt.Printf("a=%d\n", a)
		a++
	}
}
