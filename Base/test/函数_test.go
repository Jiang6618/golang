package test

import (
	"fmt"
	"testing"
)

// 函数定义
func sum(a int, b int) {
	// 定义一个求和函数
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
}

func compare(a int, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}

func TestFunc(t *testing.T) {
	// 函数调用, 无返回值
	sum(10, 20)

	// 函数调用, 有返回值
	max := compare(1, 2)
	fmt.Printf("max: %v\n", max)
}
