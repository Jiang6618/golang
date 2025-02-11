package test

import (
	"fmt"
	"testing"
)

/*
切片
是一个拥有相同类型元素的可变长度的序列, 是对数组的抽象, 是一个引用类型

可以理解为可变长度的数组, 底层是使用数组实现的,增加了自动扩容功能
切片是引用类型, 传递的是引用
*/
func TestSlice(t *testing.T) {
	// 声明切片
	/*
		与数组类似, 只要不添加长度就是声明切片
		var slice []type
		var slice []type = make([]types, len, cap)
		var slice []type = make([]types, len)
		slice := make([]type, len, cap)
		slice := make([]type, len)

		make([]T, len, cap)  // cap是可选参数, 指定切片的容量, len指定切片的长度
	*/
	var names []string
	var numbers []int

	fmt.Printf("names: %v\n", names)
	fmt.Printf("nunbers: %v\n", numbers)

	fmt.Printf("(names == nil): %v\n", (names == nil))
	fmt.Printf("(numbers == nil): %v\n", (numbers == nil))
}

func TestSlice2(t *testing.T) {
	// 切片的长度和容量

	/*
		切片拥有的长度和容量, 可以使用内置函数len()和cap()来获取
		len()获取切片的长度, cap()获取切片的容量
	*/
	var names = []string{"张三", "李四", "王五", "赵六", "田七"}
	var numbers = []int{3, 5, 1}

	fmt.Printf("len(names): %v\n", len(names))
	fmt.Printf("cap(names): %v\n", cap(names))

	fmt.Printf("len(numbers): %v\n", len(numbers))
	fmt.Printf("cap(numbers): %v\n", cap(numbers))

	var s1 = make([]string, 2, 3) // 声明一个长度为2, 容量为3的字符串类型切片
	fmt.Printf("len(s1): %v\n", len(s1))
	fmt.Printf("cap(s1): %v\n", cap(s1))
}

func TestSlice3(t *testing.T) {
	// 切片的初始化

	/*
		1. 直接初始化
	*/

	s := []int{1, 2, 3}
	fmt.Printf("s: %v\n", s)
	fmt.Println("=====================")

	/*
		2. 使用make()函数初始化
	*/

	s1 := make([]int, 3, 5)
	fmt.Printf("s: %v\n", s1)
	fmt.Println("=====================")
	/*
		3. 使用切片表达式初始化
	*/

	a := [5]int{1, 2, 3, 4, 5}
	s2 := a[1:3] // 从数组a中获取下标从1到3的切片
	s3 := a[2:]  //
	s4 := a[:3]  //
	s5 := a[:]   //
	fmt.Printf("s2: %v\n", s2)
	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s4: %v\n", s4)
	fmt.Printf("s5: %v\n", s5)
	fmt.Println("=====================")

	/*
		4. 空切片
		一个切片在未初始化之前默认为nil, 长度0 容量0
	*/

	var s6 []int
	fmt.Printf("(s6 == nil): %v\n", (s6 == nil))
	fmt.Printf("len(s6): %v\n", len(s6))
	fmt.Printf("cap(s6): %v\n", cap(s6))
}

func TestSlice4(t *testing.T) {
	// 切片遍历
	/*
		1. for循环
	*/

	s1 := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[i]: %v\n", s1[i])
	}
	fmt.Println("=====================")

	/*
		for range 循环
	*/

	for index, v := range s1 {
		fmt.Printf("a1[%d]: %d\n", index, v)
	}
}

/*
	切片是引用类型, 通过赋值的方式,会修改原有的内容
	在go中提供了copy()函数来拷贝切片
*/

func TestSlice5(t *testing.T) {
	// 添加元素
	/*
		使用append() 函数添加元素
	*/
	s1 := []int{}
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3, 4, 5) // 添加多个元素
	fmt.Printf("s1: %v\n", s1)

	s3 := []int{3, 4, 5}
	s4 := []int{1, 2}
	s4 = append(s4, s3...) // 添加另一个切片
	fmt.Printf("s4: %v\n", s4)
}

func TestSlice6(t *testing.T) {
	// 删除元素
	/*
		go中没有删除切片元素的专用方法, 使用切片本身的特性来实现

		公式: 要从切片a中删除索引为index的元素,操作方法是:
			a = append(a[:index], a[index:]...)
	*/
	s1 := []int{1, 2, 3, 4, 5}
	s1 = append(s1[:2], s1[3:]...)
	fmt.Printf("s1: %v\n", s1)
}

func TestSlice7(t *testing.T) {
	// 拷贝切片
	s1 := []int{1, 2, 3}
	s2 := s1 // 赋值的情况下,原来的变量被修改了
	s1[0] = 100
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("=====================")

	s3 := make([]int, 3)
	copy(s3, s1) // 使用copy函数, 原来的变量没有被修改
	s1[0] = 1

	fmt.Printf("s3: %v\n", s3)
	fmt.Printf("s1: %v\n", s1)
}
