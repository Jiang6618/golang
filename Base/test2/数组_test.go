package test

import (
	"fmt"
	"testing"
)

/*
	数组是相同数据类型的固定长度的序列
	一旦声明, 数组的长度就是固定的, 不能动态增加或减少
	可以通过下标访问数组的元素, 下标是从0开始的
	数组的长度和类型共同组成了数组的类型

	在 Go 语言中，数组（Array）不是引用类型，而是值类型。
	这意味着数组在赋值或传递时会创建一个新的副本，而不是引用原数组的内存地址

	当数组作为函数参数传递时，函数接收的是数组的副本，而不是原数组
*/

func TestArray(t *testing.T) {
	// 数组的定义
	/*
		数组的长度和类型共同组成了数组的类型
	*/
	var a [3]int    // 定义了一个长度为3的int类型数组
	var s [2]string // 定义了一个长度为2的string类型数组
	var b [2]bool   // 定义了一个长度为2的bool类型数组
	fmt.Printf("a: %T\n", a)
	fmt.Printf("s: %T\n", s)
	fmt.Printf("b: %T\n", b)
}

func TestArray1(t *testing.T) {
	// 没有初始化的数组
	/*
		初始化就是给数组元素赋值, 没有初始化的数组
		默认元素都是零值\布尔值为false\字符串为空字符串
	*/
	var a [3]int    // 定义了一个长度为3的int类型数组
	var s [2]string // 定义了一个长度为2的string类型数组
	var b [2]bool   // 定义了一个长度为2的bool类型数组
	fmt.Printf("a: %v\n", a)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("b: %v\n", b)
}

func TestArray2(t *testing.T) {
	// 使用初始化列表
	/*
		使用初始化列表就是将值写在大括号里面, 用逗号分隔
	*/
	var a = [3]int{1, 2, 3}
	var s = [2]string{"hello", "world"}
	var b = [2]bool{true, false}

	a1 := [3]int{1, 2, 3} // 推断类型

	fmt.Printf("a: %v\n", a)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a1: %v\n", a1)
}

func TestArray3(t *testing.T) {
	// 省略数组长度
	/*
		数组长度可以省略不写, 使用...代替, Go语言会根据初始化列表的元素个数来推断数组的长度
	*/

	var a = [...]int{1, 2, 3}
	var s = [...]string{"hello", "world"}
	var b = [...]bool{true, false}

	a1 := [...]int{1, 2, 3} // 推断类型

	fmt.Printf("a: %v\n", a)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a1: %v\n", a1)
}

func TestArray4(t *testing.T) {
	// 指定索引值的方式来初始化
	/*
		可以通过指定索引值的方式来初始化数组,未指定索引的元素将用零值初始化
	*/

	var a = [5]int{2: 10, 3: 20}
	var s = [5]string{3: "hello", 4: "world"}
	var b = [5]bool{0: true, 4: false}

	a1 := [5]int{2: 10, 3: 20} // 推断类型

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("s: %v\n", s)
	fmt.Printf("a1: %v\n", a1)
}

func TestArray5(t *testing.T) {
	// 访问和修改数组元素
	/*
		通过下标访问和修改数组元素, 下标是从0开始的, 最大下标为数组长度-1, 大于最大下标会引发越界错误
	*/

	var a [2]int
	a[0] = 100
	a[1] = 200

	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Printf("a[1]: %v\n", a[1])

	// 修改 a[0] a[1]
	a[0] = 300
	a[1] = 400
	fmt.Printf("a[0]: %v\n", a[0])
	fmt.Printf("a[1]: %v\n", a[1])

	// a[2] = 500 // 越界 error: invalid array index 2 (out of bounds for 2-element array)
	// fmt.Printf("a[2]: %v\n", a[2])  // 越界 error: invalid array index 2 (out of bounds for 2-element array)
}

func TestArray6(t *testing.T) {
	// 数组的遍历
	/*
		1. 根据数组长度,使用for循环遍历数组, 数组长度使用len()函数获取
	*/

	var a = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	/*  2. 使用range遍历数组 */
	for i, v := range a {
		fmt.Printf("a[%d]: %d\n", i, v)
	}
}
