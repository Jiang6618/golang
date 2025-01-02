package main

import "fmt"

func defineVariable() {
	// 单个声明
	var name string
	name = "jiang"

	//多个声明
	var (
		i     int64
		age   uint8
		money float64
	)

	//声明并赋值
	var val int32 = 12

	//快捷声明
	key := "apple"

	fmt.Printf("name: %s, age: %d, money: %.2f\n", name, age, money)
	fmt.Println(i, val, key)
}

func main() {
	fmt.Println("Hello World")
	fmt.Print("++++++++++++++++++++++++++++++++++++\n")

	defineVariable()
}
