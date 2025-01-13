package test

import (
	"fmt"
	"testing"
)

func TestConst(t *testing.T) {
	//声明单个常量
	const e = 2.7

	//批量声明常量
	const (
		c1 = "string"
		c2 = 22
		c3 // 继承c2
	)
	fmt.Println(e, c1, c2, c3) // 2.7 string 22 22

	//Go语⾔中的常量是⽆类型的，
	//这意味着常量在赋值给变量或表达式时，会根据上下⽂的需要⾃动进⾏类型转换
	var a int = 22
	if a == c2 {
		fmt.Println("ok")
	}

}

func TestIota(t *testing.T) {
	//iota 是 Go 中的常量生成器，从 0 开始递增，每声明一个新常量，iota 的值加 1
	const (
		l1 = iota
		l2 = iota
		l3 = iota
	)

	const (
		_ = iota
		n1
		n2
		n3
	)

	const (
		m0 = iota
		_
		m2
		m3
	)
	fmt.Println(l1, l2, l3) //0 1 2
	fmt.Println(n1, n2, n3) //1 2 3
	fmt.Println(m0, m2, m3) //0 2 3
}

func TestIota1(t *testing.T) {
	/*
		关于 iota
			每个 const 块中，iota 从 0 开始递增。
			如果一行定义了多个常量（如 a, b），它们共享同一个 iota 值。
			在位移操作和多值赋值中，iota 是一个非常简洁和强大的工具
	*/
	const (
		//1 << (10 * iota) 表示将数字 1 左移 10 * iota 位。
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
	)

	const (
		a, b = iota, iota + 2
		c, d = iota, iota + 2
		e, f = iota, iota + 2
	)

	fmt.Println(KB, MB, GB)       // 1024 1048576 1073741824
	fmt.Println(a, b, c, d, e, f) // 0 2 1 3 2 4
}
