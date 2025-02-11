package test

import (
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	// 布尔类型
	var b1 bool = true
	var b2 bool = false
	var b3 = true
	var b4 = false
	b5 := true
	b6 := false

	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("b2: %v\n", b2)
	fmt.Printf("b3: %v\n", b3)
	fmt.Printf("b4: %v\n", b4)
	fmt.Printf("b5: %v\n", b5)
	fmt.Printf("b6: %v\n", b6)
}

func TestBool1(t *testing.T) {
	// 用在条件判断中
	age := 18
	ok := age >= 18
	if ok {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}
}

func TestBool2(t *testing.T) {
	// 用在循环语句中
	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i: %v\n", i)
	}
}

func TestBool3(t *testing.T) {
	// 用在逻辑表达式中
	age := 18
	gender := "男"

	if age >= 18 && gender == "男" {
		fmt.Println("成年男性")
	} else {
		fmt.Println("其它")
	}
}
