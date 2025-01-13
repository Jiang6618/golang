package test

import (
	"fmt"
	"strings"
	"testing"
)

func outer(num int) func() {
	count := 0
	inner := func() {
		count += num
		fmt.Println(count)
	}
	return inner
}

func TestClosuer(t *testing.T) {
	counter := outer(12)
	counter()
	counter()
	counter()
}

// 闭包例子
func makeSuffixFunc(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func TestSuffix(t *testing.T) {
	addJpg := makeSuffixFunc(".jpg")
	fmt.Println(addJpg("test"))
	fmt.Println(addJpg("test.jpg"))
}
