package test

import (
	"fmt"
	"testing"
)

/*
	闭包
		可以理解成定义在一个函数内部的函灵敏
		本质上, 闭包是将函数内部和函数外部连接起来的桥梁
		指的是一个函数和与其相关的引用环境组合而成的实体

		简单来说: 闭包=函数+引用环境
*/

func addByclose1() func(a int) int {
	// 返回一个函数
	var x int
	return func(a int) int {
		x += a
		return x
	}
}

func TestClose1(t *testing.T) {
	/*
		示例中, 变量f是一个函数并且它收用了其外部作用域中的x变量
		此时f就是一个闭包

		在f的生命周期内, 变量x也一直有效
	*/
	f := addByclose1()
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(1): %v\n", f(1))
	fmt.Printf("f(20): %v\n", f(20))
	fmt.Println("============")

	f1 := addByclose1() // f1也是一个闭包, f与f1是两个闭包, 各自中的x是独立且互不影响的
	fmt.Printf("f1(20): %v\n", f1(20))
	fmt.Printf("f1(11): %v\n", f1(11))
	fmt.Printf("f1(40): %v\n", f1(40))
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func addByclose2(base int) func(a int) int {
	return func(a int) int {
		base += a
		return base
	}
}
func TestClose2(t *testing.T) {
	/*
		自定义add方法的初始值
	*/

	f := addByclose2(100)
	fmt.Printf("f(10): %v\n", f(10))
	fmt.Printf("f(1): %v\n", f(1))
	fmt.Printf("f(20): %v\n", f(20))
	fmt.Println("============")

	f1 := addByclose2(200)
	fmt.Printf("f1(20): %v\n", f1(20))
	fmt.Printf("f1(11): %v\n", f1(11))
	fmt.Printf("f1(40): %v\n", f1(40))
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func retTowFunc(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}

	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub
}

func TestClose3(t *testing.T) {
	/*
		返回多个闭包函数
	*/
	add, sub := retTowFunc(100)
	r := add(10)
	fmt.Printf("r: %v\n", r)
	r = sub(9)
	fmt.Printf("r: %v\n", r)
}
