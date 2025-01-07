package test

import (
	"fmt"
	"sort"
	"testing"
)

// map的KEY必须是可比较的类型,因为KEY是唯一的.引用类型不能做为KEY

func TestMap(T *testing.T) {
	// 先声明, 后赋值 方法一
	var m1 map[int]string // 声明变量,未初始化,值为nli, 需要make初始化
	fmt.Println(m1)
	m1 = make(map[int]string) // 初始化map, 需先声明
	m1[1] = "name"            // 赋值
	fmt.Println(m1)

	// 声明同时赋值 方法二
	m2 := map[string]int{
		"name": 1,
		"age":  20,
	} // 声明变量, 可以有初始化的值, 可以直接赋值
	m2["name2"] = 2 // 直接赋值
	fmt.Println(m2)
}

func TestMap1(t *testing.T) {
	// delete 函数删除map中指定的KEY
	m := map[string]string{
		"phone": "xiaomi",
		"car":   "Tesla",
	}

	fmt.Println(m)

	delete(m, "car")
	fmt.Println(m)
}

func TestMap2(t *testing.T) {
	// 获取map的一个值时,可以接收两个值, 第一个值是value,第二个是一个判断是否存key的bool值
	// 存在就是true, 不存在就是false
	m := map[string]string{
		"phone": "xiaomi",
		"car":   "Tesla",
	}

	if val, ok := m["name"]; ok {
		fmt.Println("name", val)
	} else if _, ok = m["car"]; ok {
		fmt.Println("car found")
	} else {
		fmt.Println("not found key")
	}
}

func TestMap3(t *testing.T) {
	// map的遍历, 是没有顺序的,而是随机的
	m := map[string]string{
		"phone": "xiaomi",
		"car":   "Tesla",
	}

	for key, val := range m {
		fmt.Println(key, val)
	}
}

func TestMap4(t *testing.T) {
	// map是没有循序的, 但是可以通过将key按顺序存到切片中, 然后顺序遍历map
	list := make([]string, 0)
	m := map[string]string{
		"d": "dd",
		"c": "cc",
		"b": "bb",
		"a": "aa",
	}

	for key := range m {
		list = append(list, key)
	}

	sort.Strings(list)

	for _, val := range list {
		fmt.Println(val, m[val])
	}
}
