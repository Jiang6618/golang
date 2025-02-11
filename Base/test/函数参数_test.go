package test

import (
	"fmt"
	"testing"
)

/*
	go语言函数可以有0个或多个参数,参数需要指定数据类型
	声明函数时的参数列表叫形参, 调用时传递的参数叫做实参

	go语言通过传值的方式传参,意味着传递给函数的是拷贝后的副本
	所以函数内问访问\修改的也是这个副本

	go语言可以使用变长参数, 有不能确定参数个数时, 可以使用变长参数
		在函数定义语句中能数部分使用 ARGE ...TYPE 的方式
		...代表的参数全部保存到一个名为ARGE的slice中
		注意这些参数的数据类型都是TYPE

*/

func func1(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestFuncVal1(t *testing.T) {
	/*
		go语言传参
	*/

	r := func1(1, 2)
	fmt.Printf("r: %v\n", r)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func func2(a int) {
	a = 200
	fmt.Printf("a1: %v\n", a)
}

func TestFuncVal2(t *testing.T) {
	/*
		参数传递, 按值传递
		调用函数func1后, a的值并没有改变,说明参数是一个拷贝的副本
		即拷贝了一份新的内容进行运算

	*/
	a := 100
	func2(a)
	fmt.Printf("a: %v\n", a)
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func func3_1(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

func func3_2(name string, age int, args ...int) {
	fmt.Printf("name: %v \n", name)
	fmt.Printf("age: %v \n", age)
	for _, v := range args {
		fmt.Printf("v: %v \n", v)
	}
}

func TestFuncval3(t *testing.T) {
	/*
		变长参数
	*/
	func3_1(1, 2, 3)
	fmt.Println("----------------")

	func3_1(1, 2, 3, 4, 5, 6)
	fmt.Println("----------------")

	func3_2("Tom", 30, 1, 2, 3)

}

func func4(a []int) {
	a[0] = 100
}

func TestFunc4(t *testing.T) {
	/*
		map slice interface channel 这些数据类型本身就是指针变型的
		拷贝传值也是拷贝的指针
		拷贝后的参数仍然指定底层的数据结构,所在修改它们可能会影响外部数据结构的值
	*/
	a := []int{1, 2}
	func4(a)
	fmt.Printf("a: %v\n", a)
}
