package test

import (
	"fmt"
	"testing"
)

func TestForRange(t *testing.T) {
	//for range 可以遍历数组，切⽚，Map，channel（通道）
	list := []int{3, 6, 9}
	for i, item := range list {
		fmt.Println("index:", i, "item:", item)
	}
}

func TestFor(t *testing.T) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	var n int
	for ; n < 5; n++ {
		fmt.Println(n)
	}
	fmt.Println()

	var m int
	for m < 5 {
		fmt.Println(m)
		m++
	}
	fmt.Println()
}
