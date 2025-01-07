package test

import (
	"fmt"
	"testing"
)

func TestGoto(t *testing.T) {
	// goto 跳到某个标签处
	for i := 0; i < 10; i++ {
		if i == 6 {
			goto here
		}
		fmt.Println(i)
	}
here:
	fmt.Println("here is here")
}

func TestBreak(t *testing.T) {
	// break 跳出循环
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
}

func TestContinue(t *testing.T) {
	// continue 跳过当前循环
	for i := 0; i < 10; i++ {
		if i == 6 {
			continue
		}
		fmt.Println(i)
	}
}
