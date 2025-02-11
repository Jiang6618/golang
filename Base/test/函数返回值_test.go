package test

import (
	"fmt"
	"testing"
)

func f1() {
	// 没有返回值
	fmt.Println("没有返回值, 只是进行一些计算")
}

func sum1(a int, b int) (ret int) {
	// 有一个返回值
	/*
		当返回值有名称时, 必须使用括号包围,逗号分隔,即使只有一个返回值
	*/
	ret = a + b
	return ret
}

func sum11(a int, b int) (c int, ok bool) {
	// 多个返回值
	/*
		1. 当返回值有名称时, 必须使用括号包围,逗号分隔,即使只有一个返回值
		2. return中可以使用表达式,但不能出现赋值表达式, 例如 return a+b是正确的
		   但 return c=a+b是错误的
		3. return关键字中指定了参数时, 返回值可以不用名称
	*/

	if a > b {
		return a + b, true
	}

	return a + b, false
}

func f2() (name string, age int) {
	// 多个返回值, 且在return中指定返回的内容
	/*
		命名的返回值是预先声明好的, 在函数内部可以直接使用,无需再次声明
		命名返回值的名称不能和函数参数名称相同, 否则报错提示变量重复定义
	*/
	name = "jiang" // 直接使用返回值
	age = 30
	return name, age
}

func f21() (string, int) {
	// 多个返回值, 且在return中指定返回的内容
	/*
		return关键字中指定了参数时, 返回值可以不用名称
	*/
	name := "jiang"
	age := 21
	return name, age
}

func f3() (name string, age int) {
	// 多个返回值, 返回值名称没有被使用
	/*
		如查retrun省略参数, 则返回值部分必须带名称
	*/
	name = "jiang"
	age = 36
	return // 等价于 retrun name, age
}

func f4() (name string, age int) {
	// return 覆盖命名返回值, 返回值名称没有被使用
	/*
		即使返回值命名了, return中也可以强制指定其它返回值的名称
		也就是说return的优先级更高
	*/
	n := "江"
	a := 18
	return n, a
}

func TestFuncReturn(t *testing.T) {

	/*
		1. return 可以有参数,也可以没有参数, 返回值可以有名称,也可以没有名称
		   go中的函数可以有多个返回值

		2. return关键字中指定了参数时,返回值可以不用名称, 如果return省略了参数,
		   则返回值部分必须带名称

		3. 当返回值有名称时, 必须使用括号包围,逗号分隔,即使用只有一个返回值

		4. 即使返回值命名了,return中也可以强制指定其它返回值的名称, return的优先级更高

		5. 命名的返回值是预先声明好的,在函数内部可以直接使用,无需再次声明.
		   命名返回值名称不能和函数参数相名称相同,否则报错提示重复定义

		6. return中可以使用表达式, 但不能出现赋值表达式
	*/

	f1()
	fmt.Println("=====================")

	ret := sum1(10, 20)

	fmt.Printf("sum: %v\n", ret)
	fmt.Println("=====================")

	ret11, ok := sum11(11, 20)

	fmt.Printf("sum11: %v\n", ret11)
	fmt.Printf("\"a>b\": %v\n", ok)
	fmt.Println("=====================")

	name2, age2 := f2()

	fmt.Printf("name2: %v\n", name2)
	fmt.Printf("age2: %v\n", age2)
	fmt.Println("=====================")

	name21, age21 := f21()

	fmt.Printf("name21: %v\n", name21)
	fmt.Printf("age21: %v\n", age21)
	fmt.Println("=====================")

	name3, age3 := f3()

	fmt.Printf("name3: %v\n", name3)
	fmt.Printf("age3: %v\n", age3)
	fmt.Println("=====================")

	name4, age4 := f4()

	fmt.Printf("name4: %v\n", name4)
	fmt.Printf("age4: %v\n", age4)
	fmt.Println("=====================")

}
