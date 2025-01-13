package test

import (
	"fmt"
	"testing"
)

//对于指针，我们需记住两个符号:&(取值的内存地址)和*(根据内存地址取值)。

func TestPointer(t *testing.T) {
	var s string
	var sp *string
	s = "abc"
	sp = &s
	fmt.Println(s, sp, &s, &sp, *sp) //abc 0xc00010e330 0xc00010e330 0xc000110070 abc
}

func TestPointerCreate(t *testing.T) {
	//第⼀种⽅式，将变量的地址赋值给指针类型
	var n = 1
	np := &n
	fmt.Println(np, *np) //0xc00000a428 1

	//第⼆种⽅式，⽤关键词 new
	p := new(int) //0xc00000a430
	fmt.Println(p)
	fmt.Printf("%T", p) //*int
}

func TestPointerGetData(t *testing.T) {
	v := 10
	p := &v
	fmt.Println(p)  //0xc00000a438
	fmt.Println(*p) //10

	*p = 20
	fmt.Println(v) //20
}
