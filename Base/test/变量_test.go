package test

import (
	"fmt"
	"testing"
)

/*
	go语言中的变量
	1. 需要先声明后使用
	2. 同一作用域中不支持重复声明
	3. 变量声明后必须使用
	4. 在声明时会自动对变量对应的内存区域进行初始化操作
	5. 每个变量会被初始化成其类型的默认值,
	   整型0, 字符串""空字符串,布尔falseconst
	   切片\函数\指针变量默认为nil
	6. 短变量声明,只能在函数内部,函数外不能使用
*/

func TestVar(t *testing.T) {
	// 变量声明
	var name string
	var age int
	var ok bool
	fmt.Println(name, age, ok)
}

func TestVar1(t *testing.T) {
	// 批量声明
	var (
		name string
		age  int
		ok   bool
	)
	fmt.Println(name, age, ok)
}

func TestVar2(t *testing.T) {
	// 变量初始化
	var name string = "Jiang"
	var age int = 35
	var site string = "jiang.com"
	fmt.Println(name, age, site)
}

func TestVar3(t *testing.T) {
	// 类型推导
	var name = "Jiang"
	var site = "jiang.com"
	var age = 30
	var ok = true
	fmt.Printf("name: %v, %T\n", name, name)
	fmt.Printf("sit: %v, %T\n", site, site)
	fmt.Printf("age: %v, %T\n", age, age)
	fmt.Printf("ok: %v, %T\n", ok, ok)
}

func TestVar4(t *testing.T) {
	// 初始化多个变量
	var name, age, site, ok = "Jiang", 35, "jiang.com", true
	fmt.Printf("name: %v, %T\n", name, name)
	fmt.Printf("sit: %v, %T\n", site, site)
	fmt.Printf("age: %v, %T\n", age, age)
	fmt.Printf("ok: %v, %T\n", ok, ok)
}

func TestVar5(t *testing.T) {
	/*
		短变量声明:
		这种方法只适合在函数内部, 函数外面不能使用
	*/
	name := "Jiang"
	age := 35
	site := "jiang.com"
	ok := false
	fmt.Printf("name: %v, %T\n", name, name)
	fmt.Printf("sit: %v, %T\n", site, site)
	fmt.Printf("age: %v, %T\n", age, age)
	fmt.Printf("ok: %v, %T\n", ok, ok)
}

func TestVar6(t *testing.T) {
	/*
		匿名变量
		接收到多个变量, 有些变量用不到时, 可以使用下划线 _ 表示变量名称
		这种变量叫做匿名变量
	*/

	name, _ := getNameAndAge()
	fmt.Printf("name: %v\n", name)
}

func getNameAndAge() (string, int) {
	return "Jiang", 35
}
