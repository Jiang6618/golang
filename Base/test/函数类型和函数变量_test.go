package test

import (
	"fmt"
	"testing"
)

/*
	在 Go 语言中，函数也是一种类型，可以像其他类型（如 int、string 等）一样被使用。
	函数类型和函数变量是 Go 中非常强大的特性，允许你将函数作为值传递、赋值给变量、
	或者作为参数和返回值。

	可以使用type关键字来定义一个函数类型. 语法格式:
		type fun func(int, int) int

	上面语句定义了一个fun函数类型
	它是一种函数类型,这种函数接收两个int类型的参数,返回一个int类型的返回值

*/

// 定义一个fun函数类型
type MyFunc func(int, int) int

// 定义两个这样结构与 MyFunc 函数类型的函数sum max
func sum3(a int, b int) int {
	return a + b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func TestFuncType1(t *testing.T) {
	/*
		使用场景1: 将函数赋值给变量, 然后通过变量调用函数
	*/
	var f MyFunc // 声明一个fun类型的函数变量f

	f = sum3 // 把 sum 赋值给函数变量f
	s := f(1, 3)
	fmt.Printf("s: %v\n", s)

	f = max
	m := f(3, 4) // 把max 赋值给函数变量f
	fmt.Printf("m: %v\n", m)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func calculate(a int, b int, operation MyFunc) int {
	return operation(a, b)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func TestFuncType2(t *testing.T) {
	/*
		使用场景2: 作为参数传递给其它函数
	*/
	var f MyFunc
	f = add
	r := calculate(1, 2, f)
	fmt.Printf("r: %v\n", r)

	f = multiply
	r = calculate(1, 2, f)
	fmt.Printf("r: %v\n", r)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func getOperation(operationType string) MyFunc {
	/* 返回一个函数,类型为Myfunc */
	if operationType == "add" {
		// 返回一个函数, 且函数的必须与MyFunc的结构一致
		return func(a, b int) int {
			return a + b
		}
	} else if operationType == "multiply" {
		return func(a, b int) int {
			return a * b
		}
	}
	return nil
}

func TestFuncType3(t *testing.T) {
	/*
		使用场景3: 从函数中返回另一个函数, 实现灵活的逻辑
	*/
	// 获取 add 函数
	operation := getOperation("add")
	fmt.Printf("Addition: %v\n", operation(1, 2))

	// 获取 multiply 函数
	operation = getOperation("multiply")
	fmt.Printf("Multiplication: %v\n", operation(2, 3))
}
