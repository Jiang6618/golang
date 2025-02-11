package test

import (
	"fmt"
	"testing"
)

type user struct {
	name string
}

func TestType(t *testing.T) {
	u := user{"Jiang"}
	b := true
	num := 1024
	s := "welcom"

	fmt.Printf("%%v  - 格式化输出结构:				%v\n", u)
	fmt.Printf("%%#v - 输出值的go表示方法:			%#v\n", u)
	fmt.Printf("%%T  - 输出值类型:          			%T\n", u)

	fmt.Printf("%%t  - 输出值的true或false: 			%t\n", b)

	fmt.Printf("%%c  - 数值对应的Unicode编码字符:		%c\n", num)
	fmt.Printf("%%U  - Unicode表示:				%U\n", num)

	fmt.Printf("%%d  - 十进制表示:				%d\n", num)
	fmt.Printf("%%b  - 二进制表示:				%b\n", num)
	fmt.Printf("%%o  - 八进制表示:				%o\n", num)
	fmt.Printf("%%x  - 十六进制表示(用a-f):			%x\n", num+110)
	fmt.Printf("%%X  - 十六进制表示(用A-F):			%X\n", num+110)

	fmt.Printf("%%s  - 直接输出字符串者[]byte:			%s\n", s)
	fmt.Printf("%%q  - 双引号括起来的字符串:			%q\n", s)
}
