package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	var nums [4]int
	fmt.Println(nums)

	nums = [4]int{1, 2, 3, 4}
	fmt.Println(nums)

	var numbers = [3]int{1, 2, 3}
	fmt.Println(numbers)

	s := [2]string{"a", "b"}
	fmt.Println(s)
}

func TestChangeArray(t *testing.T) {
	nums := [4]int{1}
	fmt.Println(nums)
	nums[0] = 10
	fmt.Println(nums)
	nums[0] = 100
	fmt.Println(nums)
}

func change(arr [4]int) {
	fmt.Printf("%p\n", &arr)

}
func TestChange(t *testing.T) {
	var numbers [4]int
	numbers[0] = 1
	fmt.Printf("%p\n", &numbers)
	change(numbers)
}

func TestArrType(t *testing.T) {
	nums1 := [4]int{1, 2, 3, 4}
	nums2 := [5]int{1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(nums1), reflect.TypeOf(nums2))
}

func TestArrOmitNums(t *testing.T) {
	var nums = [...]int{1, 2, 3}
	fmt.Println(nums, reflect.TypeOf(nums))
}

func TestArrPtr(t *testing.T) {
	a := [4]*int{0: new(int), 3: new(int)}
	fmt.Println(a) // [0xc000014310 <nil> <nil> 0xc000014318]

	*a[0] = 10
	// *a[2] = 30 // 如果指针地址为空, 是会报空指针错误的, 比如panic: runtime error: invalid memory address or nil pointer
	*a[3] = 20
	fmt.Println(a)

	fmt.Println(*a[0], *a[3])

	a[2] = new(int) // 为a[2]赋值
	*a[2] = 40
	fmt.Println(*a[0], *a[2], *a[3])
}

func TestArrCopy(t *testing.T) {
	a1 := [4]string{"a", "b", "c", "d"}
	a2 := [4]string{"x", "y", "z", "h"}

	a1 = a2
	fmt.Println(a1, a2)
}

func TestArrayPtrCopy(t *testing.T) {
	a1 := [4]*string{new(string), new(string), new(string), new(string)}
	a2 := a1
	fmt.Println(a1, a2)

	*a1[0] = "A"
	*a1[1] = "B"
	*a1[2] = "C"
	*a1[3] = "C"
	fmt.Println(*a1[0], *a2[0])
}

func TestArrayIter(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func TestArray2x(t *testing.T) {
	pos := [4][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	fmt.Println(pos)

	pos[0] = [2]int{10, 10}
	fmt.Println(pos)
}

func TestArray(t *testing.T) {
	a := [2]int{}
	fmt.Println(a, a[1])
}
