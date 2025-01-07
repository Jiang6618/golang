package test

import (
	"errors"
	"fmt"
	"testing"
)

// func FunctionName (param Type, param2 Type2 ...) (Type1, Type2) {}
// func 函数名(参数1 类型, 参数2 类型, ...) (返回值1 类型, 返回值2 类型) {}

func TestFunc(t *testing.T) {
	// 函数参数的几种方式
	SayHello()
	Say("你好")
	SaySomeTimes(5, "Welcome")
}

func SayHello() {
	// 不带参数
	fmt.Println("Hello")
}

func Say(work string) {
	// 带一个参数
	fmt.Println(work)
}

func SaySomeTimes(some int, work string) {
	// 带多个参数
	for i := 0; i < some; i++ {
		fmt.Println(i, work)
	}
}

// +++++++++++++++++++++++++++++++++++++++++++++++++

func TestFunc1(t *testing.T) {
	fmt.Println(AddTwo(1, 2))
	fmt.Println(Sum(1, 2, 3, 4))
	nums := []int{4, 5, 6}
	fmt.Println(Sum2(nums...))
	fmt.Println(Check(9))
	fmt.Println(Check(11))
}

func AddTwo(a, b int) int {
	// 一个返回值
	return a + b
}

func Sum(number ...int) int {
	// 不定参数, 可以为0个, 也可以为多个
	var res int
	for _, item := range number {
		res += item
	}
	return res
}

func Sum2(number ...int) (res int) {
	// 返回值的变量可以提前定义
	for _, item := range number {
		res += item
	}
	return
}

func Check(i int) (bool, error) {
	// 返回多个值
	if i > 10 {
		return false, errors.New("number not greate than 10")
	} else {
		return true, nil
	}
}
