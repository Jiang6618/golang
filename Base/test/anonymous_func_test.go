package test

import (
	"fmt"
	"testing"
)

// 匿名函数
func TestAFunc(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
}

func TestSub(t *testing.T) {
	var sub func(int, int) int

	sub = func(a, b int) int {
		return a - b
	}
	fmt.Println(sub(25, 10))
}
