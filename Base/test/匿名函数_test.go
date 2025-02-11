package test

import (
	"fmt"
	"testing"
)

/*
	go语言函数不能嵌套
	但是可以在函数内部定义匿名函数

	匿名函数就是没有名称的函数, 语法格式:
		func (参数列表)(返回值) //可以即没有参数,也没有返回值
*/

func TestAnonymous1(t *testing.T) {
	// 定义一个匿名函数,并赋值给变量
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	i := max(1, 2) // 调用函数变量
	fmt.Printf("i: %v\n", i)
}

func TestAnonymous2(t *testing.T) {

	/*
		定义一个匿名函并执行
	*/

	func(x, y int) {
		fmt.Printf("x + y = %v\n", (x + y))
	}(1, 2)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}(10, 20)
	fmt.Printf("max: %v\n", max)
}
