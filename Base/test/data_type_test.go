package test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	var f1 float32 = 1.2
	f2 := 1.3 // 使⽤类型推断⽅式声明浮点型变量时，该变量会被推断为 float64
	fmt.Println(reflect.TypeOf(f1), reflect.TypeOf(f2))
	/*
		float32 float64
	*/
}

func TestLiteral(t *testing.T) {
	//字面量
	v1 := 0b00101101            // 0b 代表二进制
	v2 := 0o64                  // 0o 代表八进制
	v3 := 0x1a                  // 0x 代表十六进制
	v4 := 12_45                 // 数字可用_分割
	fmt.Println(v1, v2, v3, v4) //45 52 26 1245
}

func TestBool(t *testing.T) {
	var no bool // 默认为false
	ok := true
	fmt.Println(no, ok) //false true
}

func TestString(t *testing.T) {
	/*
		字符串⼀旦定义，就不可更改，只能读或复制
		Go语⾔的字符串类型是原⽣类型，字符串的值⽤双引号（""）包括，字符串使⽤UTF-8编码，
		可以包含任意 Unicode 字符
	*/
	var str string
	str2 := "你好 world"
	str3 := `tom
			jack
			martin`
	fmt.Println(str, str2, str3)

	/*
		 你好 world tom
					jack
					martin
	*/
}

func TestString2(t *testing.T) {
	s1 := "您好jim"
	s2 := "abcd"
	s3 := s1 + s2                 //两个字符串可用+来链接
	fmt.Println(len(s1), len(s2)) // 9 4 :字符串长度表示的是字符串的byte数. 一个汉字由三个byte组成
	//
	fmt.Println(s3)
	fmt.Println(s2[1])      // 98 :通过字符串下标寻找byte
	fmt.Println(s2 > "abc") // true :字符串可以比较大小
	for _, char := range s1 {
		// 可以能过range来遍历字符串的每个字符, 为Unicode字符
		fmt.Println(char, string(char))
		/*
			24744 您
			22909 好
			106 j
			105 i
			109 m
		*/
	}
}

func TestByteAndRune(t *testing.T) {
	/*
		⽤单引号（''）包括⼀个字符，默认为rune
		Go语⾔的字付有两种类型,其中byte代表 ASCII码的⼀个字符，rune代表⼀个UTF-8字符
		byte是uint8类型的别名，rune是int32类型的别名。
	*/
	var b1 = 'a'      // 默认为rune类型
	var b2 byte = 'a' // 定义为byte类型
	var b3 rune = '我' // 除了ASCII字符, 只能定义为rune

	fmt.Println(b1, reflect.TypeOf(b1)) // 97 int32
	fmt.Println(b2, reflect.TypeOf(b2)) // 97 int8
	fmt.Println(b3, reflect.TypeOf(b3)) // 97 int32
}

func TestTypeConversion(t *testing.T) {
	/*
		类型转换
		Go语⾔只有强制类型转换，没有隐式类型转换
	*/

	// 类型转换的几种方式
	var a int = 12
	fmt.Printf("%d", a)          // 数字转字符串
	fmt.Println(strconv.Itoa(a)) // 数字转字符串
	fmt.Println(float64(a))      // 数字转浮点类型

	list := []byte{'a', 'b', 'c'}
	fmt.Println(string(list)) // 转字符串

	var b int64 = 257
	var c int8 = 34
	fmt.Println(uint8(b), int64(c)) //小类型转大类型, 会被截取

	var d = "12"
	fmt.Println(strconv.Atoi(d)) // 字符串转数字
}
