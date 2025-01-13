package test

import (
	"fmt"
	"testing"
)

/*
	结构体方法:
		1. 结构体方法是一种特殊类型的函数，它与结构体绑定在一起，可以通过结构体对象调用
		2. 结构体方法不像传统的类(Class)的概念, 它是使用接收者(receiver)的方式来绑定方法和对象.
		   类似于其它语言的this或self关键字
		3. 结构体方法的接收者可以是值类型或指针类型
*/

// 结构体方法声明
type Student2 struct {
	Id   int
	Name string
}

func (s Student2) Say() {
	fmt.Printf("%s say %s\n", s.Name, "hello")
}

func (s Student2) getName() string {
	return s.Name
}

func NewStudent(id int, name string) Student2 {
	return Student2{
		Id:   id,
		Name: name,
	}
}

func TestStudentMethod(t *testing.T) {
	student1 := NewStudent(1, "Tom")
	student1.Say()
	fmt.Println(student1.getName())
}

// 值接收者和指 针接收者
/*
	1. 值接收者
		1.1. 使用值接收者声明的方法, 在调用时会将接收者对象的一个副本传递给方法
		1.2. 值接收者声明的方法, 只能修改接收者对象的副本, 无法修改接收者对象本身
		1.3. 值接收者声明的方法, 适用于接收者对象是一个结构体实例的情况

		func (v T) MethodName(){
			// 方法体
		}

	2. 指针接收者
		2.1. 使用指针接收者声明的方法, 在调用时会将接收者对象的地址传递给方法
		2.2. 指针接收者声明的方法, 可以修改接收者对象本身
		2.3. 指针接收者声明的方法, 适用于接收者对象是一个结构体指针的情况

		func (v *T) MethodName(){
			// 方法体
		}

	在实际编码中, 不会在同⼀个类型的⽅法中混合使⽤值接收者和指针接收者
	可以根据以下规则判断应该使⽤值接收者，还是指针接收者:
	- 在⽅法中修改接收者的值时应使⽤指针接收者。
	- 当接收者是⼤对象(拷⻉代价⽐较⼤)时，应使⽤指针接收者
	- 保证⼀致性。如果⼀个类型的某个⽅法使⽤了指针接收者，那么该类型的其他的⽅法使⽤指针接收者
*/

type Student3 struct {
	Id   int
	name string
}

func (s Student3) Say() {
	fmt.Printf("%s say %s\n", s.name, "hello")
}

func (s Student3) GetName() string {
	return s.name
}

func (s *Student3) SetName(name string) {
	s.name = name
}

func TestStudent3(t *testing.T) {
	s3 := Student3{Id: 1, name: "Tom"}
	s3.Say()
	fmt.Println(s3.GetName())
	fmt.Printf("s3: %#v\n", s3)

	s3.SetName("Jerry")
	s3.Say()
	fmt.Println(s3.GetName())
	fmt.Printf("s3: %#v\n", s3)
}

// 方法结合

type Persons struct {
	Name string
	Age  int
}

func (p *Persons) Run() {
	fmt.Println(p.Name, "is running")
}

type Student4 struct {
	Persons
	Sid int
}

func (s Student4) Learn() {
	fmt.Println(s.Name, "is learning")
}

func TestRun(t *testing.T) {
	tom := Student4{
		Persons: Persons{
			Name: "Tom",
			Age:  18,
		},
		Sid: 36,
	}
	tom.Run()
	tom.Learn()
}
