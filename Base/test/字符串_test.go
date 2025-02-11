package test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// go语言中的字符串是一个任意字节的常量序列 - []byte

func TestStr(t *testing.T) {
	// 字符串字面量
	var str1 string = "hello world"
	var html string = `
		<html>
			<head><title>hello world</title></head>
		<html>
	`

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("html: %v\n", html)
}

func TestStr1(t *testing.T) {
	// 字符串连接 - 使用加号
	/*
		go中字符串都是不可变的,每次运算都会产生一个新的字符串
		生产生很多临时的无用的字符串,不仅没为, 还会给gc带来负担
		性能比较差
	*/
	name := "tom"
	age := "20"
	msg := name + " " + age
	fmt.Printf("msg: %v\n", msg)
	fmt.Println("----------------")

	msg = ""
	msg += name
	msg += " "
	msg += age
	fmt.Printf("msg: %v\n", msg)
}

func TestStr2(t *testing.T) {
	// 字符串连接 - 使用fmt.Sprintf()函数
	/*
		内部使用[]byte实现,不像直接运算符那样会产生很多临时的字符串,
		但内部逻辑比较复杂,有很多额外的判断,还用于了interface,
		性能不也是很好
	*/
	name := "tom"
	age := "20"

	msg := fmt.Sprintf("%s, %s", name, age)
	fmt.Printf("msg: %v\n", msg)
}

func TestStr3(t *testing.T) {
	// 字符串连接 - 使用 strings.Join()
	/*
		join会先根据字符串数组的内容, 计算出一个拼接之后的长度,然后申请对应大小的内存
		一个一个字符串填入,在已有一个数组的情况下效率很高, 但是本来没有,去构造这个数据
		的代价也不小
	*/
	name := "tom"
	age := "20"

	msg := strings.Join([]string{name, age}, ",")
	fmt.Printf("msg: %v\n", msg)
}

func TestStr4(t *testing.T) {
	// 字符中拼接 - 使用 buffer.WriteString()
	/*
		比较理想, 可以当成可变字符使用,对内存的增长也有优化,如果能预估字符串长度,可以用
		buffer.Grow()接口设置capacity
	*/
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString(", ")
	buffer.WriteString("20")
	fmt.Printf("buffer.String(): %v\n", buffer.String())
}

func TestStr5(t *testing.T) {
	// 字符串转义
	/*
		\r 回车符(返回行首)
		\n 换行符(直接跳到下一行的同列位置)
		\t 制表符
		\' 单引号
		\" 双引号
		\\ 反斜杠
	*/

	fmt.Print("hello\tworld\n")

	fmt.Print("\"c:\\test\\\"")
}

func TestStr6(t *testing.T) {
	// 字符中切片操作
	str := "hello world"
	n := 3
	m := 5

	fmt.Printf("str[n]: %v\n", str[n])
	fmt.Printf("str[n]: %c\n", str[n])
	fmt.Printf("str[n:m]: %v\n", str[n:m])
	fmt.Printf("str[n:]: %v\n", str[n:])
	fmt.Printf("str[:m]: %v\n", str[:m])
	fmt.Printf("str[:]: %v\n", str[:])
}

func TestStr7(t *testing.T) {
	// 字符串常用方法
	s := "Hello World"

	// len(str) 求长度
	fmt.Printf("len(s): %v\n", len(s))

	// strings.Split()  分割
	fmt.Printf("strings.Split(s, \" \"): %v\n", strings.Split(s, " "))

	// strings.Contains() 判断是否包含
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))
	fmt.Printf("strings.Contains(s, \"Hello\"): %v\n", strings.Contains(s, "Hello"))

	// strings.HasPrefix() 前缀判断
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello"))
	fmt.Printf("strings.HasPrefix(s, \"Hell\"): %v\n", strings.HasPrefix(s, "Hell"))

	// strings.HasSuffix() 后缀判断
	fmt.Printf("strings.HasSuffix(s, \"world\"): %v\n", strings.HasSuffix(s, "world"))
	fmt.Printf("strings.HasSuffix(s, \"World\"): %v\n", strings.HasSuffix(s, "World"))

	// strings.Index() | strings.LastIndex() 子串出现的位置
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))

}
