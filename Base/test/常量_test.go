package test

import (
	"fmt"
	"testing"
)

/*
	常量, 是在程序编译阶段就确定下来的值, 而程序在运行时则无法改变该值
	go中的常量可以是数值类型\布尔类型\字符串类型等

	iota 可以被认为是一个可被编译器修改的常量
		 它默认开始值是0,每调用一次加1
		 遇到const关键字时被重置为0
*/

func TestConst(t *testing.T) {

	// 定义常量的方式
	// const constantName [type] = value

	const PI float64 = 3.14

	const PI2 = 3.1415 // 可以省略类型

	const (
		width  = 100
		height = 200
	)

	const i, j = 1, 2
	const a, b, c = 1, 2, "foo"

	fmt.Println(PI, PI2, width, height, i, j, a, b, c)
}

func TestConst1(t *testing.T) {
	// const同时声明多个常量时, 如果省略了值则表示和上面一行的值相同
	const (
		a1 = 100
		a2
		a3
		a4 = 200
		a5
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)
}

func TestIota(t *testing.T) {
	/*
		iota 可以被认为是一个可被编译器修改的常量
		它默认开始值是0,每调用一次加1
		遇到const关键字时被重置为0
	*/

	const (
		a1 = iota
		a2 = iota
		a3 = iota
	)

	const a4 = 100
	const (
		a5 = iota
		a6 = iota
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
	fmt.Printf("a4: %v\n", a4)
	fmt.Printf("a5: %v\n", a5)
	fmt.Printf("a6: %v\n", a6)
}

func TestIota2(t *testing.T) {
	// 使用_跳过某些值
	const (
		a1 = iota
		_
		a2 = iota
	)

	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
}

func TestIota3(t *testing.T) {
	// iota 声明中间插队
	const (
		a1 = iota
		a2 = 100
		a3 = iota
	)
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}
