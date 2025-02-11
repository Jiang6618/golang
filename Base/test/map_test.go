package test

import (
	"fmt"
	"testing"
)

/*
	map是一种key:value键值对的数据结构容器.
	map内部实现是哈希表(hash)

	map最重要的一点是通过key来快速检索数据, key类似于索引, 指向数据的值

	map是引用类型
*/

func TestMap(t *testing.T) {
	// map 语法格式
	/*
		可以使用内建函数make, 也可以使用map关键字来定义map

		var map_variable map[key_data_type]value_data_type
		map_variable = make(map[key_data-type]value_data_type)
	*/

	var m1 = make(map[string]string)
	fmt.Printf("(m1 == nil): %v\n", (m1 == nil))

	m1["name"] = "Tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	fmt.Printf("m1: %v\n", m1)

	m2 := map[string]string{
		"name":  "kite",
		"age":   "30",
		"email": "kite@163.com",
	}
	fmt.Printf("m2: %v\n", m2)
}

func TestMap1(t *testing.T) {
	// 访问map
	/*
		通过下标key获取其值
	*/
	m1 := make(map[string]string)
	m1["name"] = "Tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	name := m1["name"]
	age := m1["age"]
	email := m1["email"]

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("email: %v\n", email)
}

func TestMap2(t *testing.T) {
	// 判断某个键是否存在
	/*
		通过 value, ok := map[key] 的方式判断键是否存在
		ok为true,存在, 否则,不存在
	*/
	var m1 map[string]string = map[string]string{
		"name":  "Tom",
		"age":   "20",
		"email": "tom@gmail.com",
	}

	iskey := "address"

	v, ok := m1[iskey]
	if ok {
		fmt.Printf("key %v存在,值为%v\n", iskey, v)
	} else {
		fmt.Printf("key %v不存在", iskey)
	}
}

func TestMap3(t *testing.T) {
	// 使用for range 循环进行map遍历, 得到key和value值

	m := make(map[string]string)
	m["name"] = "Tom"
	m["age"] = "20"
	m["email"] = "tom@gmail.com"

	// 遍历KEY
	for key := range m {
		fmt.Printf("m[%v]: %v\n", key, m[key])
	}
	fmt.Println("=====================")

	// 遍历key和value
	for key, value := range m {
		fmt.Printf("%v: %v\n", key, value)
	}
}
