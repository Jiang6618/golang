package test

import (
	"fmt"
	"testing"
)

/*

注意事项
defer 执行顺序：
	多个 defer 按照声明顺序的反序执行(先进后出)，即后声明的先执行。
panic 与 recover 的作用域：
	recover 只能捕获同一调用栈中的 panic，无法跨协程捕获。
适用场景：
	defer：用于清理资源（文件关闭、锁释放等）。
	panic：用于处理程序中不可恢复的错误。
	recover：用于捕获并处理 panic，避免程序崩溃。

*/

func riskyOperation() {
	defer func() {
		if res := recover(); res != nil {
			fmt.Println("Panic content: ", res)
		}
	}()

	fmt.Println("Executing risky operation")
	panic("Something went wrong")
	fmt.Println("operation completed") // This will not execute
}

func TestRecover1(t *testing.T) {
	fmt.Println("Start main")
	riskyOperation()
	fmt.Println("End main")
}
