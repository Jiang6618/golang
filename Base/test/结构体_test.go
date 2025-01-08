package test

import (
	"fmt"
	"testing"
)

/*
	结构体:
		Go语言提供了一种自定义数据类型, 可以将零个或多个任意类型的变量,组合在一起聚合成一个新的数据类型,
		这种数据类型叫做结构体,英文名称struct
		目地是总结归纳出事物的相同的属性,并把它们划分为某一类或者某一类型

	结构体定义:
		使用type和struct关键字定义结构体
		type 结构体名 struct {
			字段名 字段类型
			字段名 字段类型
			...
		}

		类型名称表示⾃定义结构体的名称，在同⼀个包内不能重复。
		字段名称表示结构体字段名，同⼀结构体中的字段名必须唯⼀。
		字段类型表示结构体字段的具体类型
*/

// 结构体定义
type User struct {
	Id     int    // 用户ID
	Name   string // 用户名
	Email  string // 用户邮箱
	Del    bool   // 用户是否被删除
	Parent *User  // 父级用户, 用指针类型
}

// 结构体变量

func TestStructVal(t *testing.T) {
	// 第一种方式
	var u User
	u.Id = 100
	u.Name = "张三"
	fmt.Printf("u: %#v, %p\n", u, &u)

	// 第二种方式
	u2 := User{
		Id:     200,
		Name:   "李四",
		Email:  "1111@163.com",
		Del:    false,
		Parent: &u,
	}

	fmt.Printf("u2: %#v, %#v\n", u2, u2.Parent)

	// 第三种方式
	u3 := User{300, "王五", "300@qq.com", true, nil}

	fmt.Printf("u3: %#v\n", u3)
}

// 结构体组合
/*
在 Go 语⾔中,结构体组合是⼀种⾮常强⼤和灵活的特性。
它允许我们通过将⼀个或多个结构体嵌⼊到另⼀个结构体中来创建新的结构体类型。
这种机制可以帮助我们实现代码重⽤、组合和继承等⾯向对象的特性
*/

type Persion struct {
	Name string
	Age  int
}

type Student struct {
	Persion
	Sid int
}

func TestStudent(t *testing.T) {
	tom := Student{
		Persion: Persion{
			Name: "Tom",
			Age:  18,
		},
		Sid: 100,
	}

	fmt.Printf("tom.Name: %v\n", tom.Name)
	fmt.Printf("tom: %#v\n", tom)
}
