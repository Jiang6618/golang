package test

import (
	"fmt"
	"testing"
)

// func Print(v string) {
// 	fmt.Println(v)
// }

func TestPanic(t *testing.T) {
	defer Print("关闭")
	fmt.Println("开始")
	panic("退出")
	fmt.Println("结束") // 不会执行
}
