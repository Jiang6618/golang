package test

import (
	"fmt"
	"testing"
)

func TestOperator(t *testing.T) {
	// 算术运算符
	/*
		++(自增) --(自减)在go中是单独语句, 并不是运算符
	*/
	a := 100
	b := 10

	fmt.Printf("(a + b): %v\n", (a + b))
	fmt.Printf("(a - b): %v\n", (a - b))
	fmt.Printf("(a * b): %v\n", (a * b))
	fmt.Printf("(a / b): %v\n", (a / b))
	fmt.Printf("(a %% b): %v\n", (a % b))
}

func TestOperator1(t *testing.T) {
	// 关系运算符
	a := 1
	b := 2
	fmt.Printf("(a > b): %v\n", (a > b))
	fmt.Printf("(a < b): %v\n", (a < b))
	fmt.Printf("(a >= b): %v\n", (a >= b))
	fmt.Printf("(a <= b): %v\n", (a <= b))
	fmt.Printf("(a == b): %v\n", (a == b))
	fmt.Printf("(a != b): %v\n", (a != b))
}

func TestOperator2(t *testing.T) {
	// 逻辑运算符
	a := true
	b := false

	fmt.Printf("(a && b): %v\n", (a && b))
	fmt.Printf("(a || b): %v\n", (a || b))
	fmt.Printf("(!a): %v\n", (!a))
	fmt.Printf("(!b): %v\n", (!b))
}

func TestOperator3(t *testing.T) {
	a := 4 // 二进制100
	fmt.Printf("a: %d %b\n", a, a)
	b := 8 // 二进制 1000
	fmt.Printf("b: %d %b\n", b, b)

	fmt.Printf("(a & b): %v %b\n", (a & b), (a & b))
	fmt.Printf("(a | b): %v %b\n", (a | b), (a | b))
	fmt.Printf("(a ^ b): %v %b\n", (a ^ b), (a ^ b))
	fmt.Printf("(a << b): %v %b\n", (a << b), (a << b))
	fmt.Printf("(a >> b): %v %b\n", (a >> b), (a >> b))
}

func TestOperator4(t *testing.T) {
	// 赋值运算符
	var a int
	a = 100
	fmt.Printf("a: %v\n", a)
	a += 1
	fmt.Printf("a: %v\n", a)
	a -= 1
	fmt.Printf("a: %v\n", a)
	a *= 2
	fmt.Printf("a: %v\n", a)
	a /= 2
	fmt.Printf("a: %v\n", a)
}
