package test

import (
	"fmt"
	"testing"
)

/*
recover() 函数用于终止错误处理流程。一般情况下，recover() 会在 defer 函数中使用
通常与panic搭配使用，用于捕获 panic 引起的错误，recover() 会返回 panic 的错误信息
*/
func TestRecover(t *testing.T) {
	defer func() {
		if res := recover(); res != nil {
			fmt.Println("panic内容: ", res)
		}
	}()

	fmt.Println("开始")
	panic("退出")
	fmt.Println("结束") // 不会执行
}
