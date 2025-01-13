package test

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	/*
		new 函数：new ⽤于创建某种类型的指针，并为其分配内存空间。
		它接受⼀个类型作为参数，并返回该类型的指针。被分配的内存会被初始化为零值（zero value）
	*/
	v := new(int)
	fmt.Println(v, *v)

	/*
		make 函数：make ⽤于创建并初始化引⽤类型（如切⽚、映射、通道）的数据结构。
		它接受⼀个类型、⻓度等参数，并返回该类型的初始化后的值。
	*/
	s := make([]int, 0)
	fmt.Println(s)

	/*
		new 和 make 返回的都是指针类型。
		使⽤ new 创建的指针指向的是零值，
		使⽤ make 创建的引⽤类型会进⾏初始化
	*/
}
