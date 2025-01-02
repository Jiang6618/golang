package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVar1(t *testing.T) {
	// 单个声明
	var name string

	//多个声明
	var (
		i       int64
		age     uint8
		money   float64
		isAdmin bool
	)

	//声明并赋值
	var val int32 = 12

	//快捷声明
	key := "apple"

	/*
		%d：用于整型（如 id 和 age）。
		%s：用于字符串（如 name）。
		%.2f：用于浮点型，保留两位小数（如 money）。
		%t：用于布尔值（如 isAdmin）。
	*/
	fmt.Printf("id: %d, name: %s, age: %d, money: %.2f, isAdmin: %t \n", i, name, age, money, isAdmin)
	/*
		整型和浮点型变量的默认值为 0
		字符串变量的默认值为空字符串
		布尔型变量的默认值为 false。
		切⽚、函数、指针变量的默认值为 nil。
	*/

	fmt.Println(val, key)
}

func TestVar2(t *testing.T) {
	//短声明, 只能用在函数中, 而全局变量要用var来声明
	//使用 := 声明的变量的类型,由编译器推断
	str := "dog"
	num := 100
	flt := 9.99
	ok := true

	fmt.Println(reflect.TypeOf(str)) //string
	fmt.Println(reflect.TypeOf(num)) //int
	fmt.Println(reflect.TypeOf(flt)) //float64
	fmt.Println(reflect.TypeOf(ok))  // bool
}

func TestVar3(t *testing.T) {
	//在使⽤多重赋值时，如果要忽略某个值，那么可以使⽤匿名变量(anonymous variable)。
	//匿名变量⽤下划线 _ 个表示
	//注意：匿名变量不占⽤命名空间，不会分配内存，所以不存在重复声明。
	name, _ := getUser()
	_, age := getUser()
	fmt.Println(name, age)
}

func getUser() (string, int) {
	return "Tom", 20
}
