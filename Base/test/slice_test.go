package test

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s1 := []string{"a", "c", "d", "e", "f", "h"}
	s2 := s1[1:3]
	s3 := s1[3:5]
	fmt.Println(s1, s2, s3)

	s1[1] = "xy"
	s1[4] = "yy"

	fmt.Println(s1, s2, s3)
}

func TestSlice2(t *testing.T) {
	// 切片的定义方式
	// 1. 直接定义
	var s1 []byte
	fmt.Printf("%T\n", s1)
	fmt.Println(len(s1), cap(s1))

	// 2.直接赋值
	s2 := []int{1, 2, 3}
	fmt.Println(len(s2), cap(s2))

	// 截取方式
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s3 := s[1:4]
	s4 := s[:3]
	s5 := s[5:]
	s6 := s[:]
	fmt.Println(s3, s4, s5, s6)

	// make创建
	s7 := make([]int, 2, 3)
	fmt.Println(s7, len(s7), cap(s7))
}
