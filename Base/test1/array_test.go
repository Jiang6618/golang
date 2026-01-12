package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	var arr1 [3]int
	arr2 := [3]uint{1, 2, 3}
	arr3 := [...]string{"aa", "bb"}
	fmt.Println(arr1, arr2, arr3)
}

func TestArray2(t *testing.T) {
	arr := [5]int{1, 2, 3, 4}
	fmt.Println(arr[1], arr[4], len(arr), cap(arr))
	fmt.Println(len(arr))

	arr[2] = 10
	fmt.Println(arr)

	for i, item := range arr {
		fmt.Println(i, item)
	}
}

func TestArray3(t *testing.T) {
	two := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(two[1][1])
	two[1][2] = 10
	fmt.Println(two)
}

func TestArray4(t *testing.T) {
	arr := [...]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arr4 := arr
	fmt.Printf("arr: %p. arr4: %p\n", &arr, &arr4)
	fmt.Println(reflect.TypeOf(arr), reflect.TypeOf(arr4))

	arr1 := append([]int{}, arr[1:7]...) // 坑
	arr2 := append([]int{}, arr[5:]...)

	fmt.Printf("arr: %T, arr1: %T, arr2: %T\n", arr, arr1, arr2)
	fmt.Println(reflect.TypeOf(arr), reflect.TypeOf(arr1), reflect.TypeOf(arr2))
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(arr1, len(arr1), cap(arr1))
	fmt.Println(arr2, len(arr2), cap(arr2))

	//底层数组arr的值变了, 以它为基准的切片的值也会随之变化
	arr[5] = 100
	fmt.Println(arr1, arr2)
	fmt.Println(arr, arr4)

	arr1 = append(arr1, 1, 3, 4)
	fmt.Println(arr1, arr)
}
