package test

import (
	"fmt"
	"testing"
)

func deferFunc() {
	defer Print("结束")
	fmt.Println("开始")
}

func deferFunc1() {
	defer Print("结束1")
	defer Print("结束2")
	defer Print("结束3")

	fmt.Println("开始")
}

func Print(str string) {
	fmt.Println(str)
}

func TestDefer(t *testing.T) {
	deferFunc()
	deferFunc1()
}
