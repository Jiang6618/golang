package test

import (
	"fmt"
	"testing"
)

/*
	go语言函数,
		可以作函数的参数传递给另一个函数
		也可以作为一个函数的返回値返回
*/

func sayHello(name string) {
	fmt.Println("Hello,", name)
}

func funcAsVal(name string, f func(string)) {
	f(name)
}

func TestHigherFunc1(t *testing.T) {
	/* 函数做为参数 */
	funcAsVal("Tom", sayHello)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func add1(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func cal(s string) func(int, int) int {
	switch s {
	case "+":
		return add1
	case "-":
		return sub
	default:
		return nil
	}
}

func TestHigherFunc2(t *testing.T) {
	/*
		函数作为返回值
	*/
	f := cal("+")
	fmt.Printf("res: %v\n", f(2, 2))

	f = cal("-")
	fmt.Printf("res: %v\n", f(2, 2))
}
