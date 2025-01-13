package test

import (
	"fmt"
	"testing"
)

/*
类型定义: 使用type关键字基于已存在的类型,采用类型定义的方式创建新类型
	type NewType SourceType

*/

/*
	type NewInt int                  // 定义一个新类型NewInt,基于int类型
	type ListStr []string            // 定义一个新类型ListStr,基于[]string类型
	type report map[string]int       // 定义一个新类型report,基于map[string]int类型
	type Notify func(string, report) // 定义一个新类型Notify,基于func(string, report)类型
*/

/*
类型别名: 使用type关键字基于已存在的类型,采用类型别名的方式创建新类型
	type NewType = SourceType
*/

// 示例:

func TestType(t *testing.T) {
	type NewInt int  // 使用类型定义, 定义一个新类型NewInt,基于int类型
	type MyInt = int // 使用类型别名,定义一个新类型MyInt,基于int类型
	var a int        // 定义一个int类型变量
	var b MyInt      // 定义一个MyInt类型变量
	var c NewInt     // 定义一个NewInt类型变量

	//类型别名和其源类型本质上属于同⼀个类型，它们的变量之间⽀持直接赋值，⽆须进⾏类换
	a = b // 类型相同,可以赋值

	// ⽤类型定义的⽅式定义了全新的类型，该类型与源类型的底层类型相同，⽀持显式转换。
	// 但是不同类型之间不能直接赋值
	// c = b // 类型不相同,不可以赋值
	c = NewInt(b) // 类型不同,需要强制类型转换

	fmt.Println(a, b, c)
	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)
}
