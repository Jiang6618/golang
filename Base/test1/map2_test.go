package test

import (
	"fmt"
	"testing"
)

func TestIMap(t *testing.T) {
	var m map[string]int
	fmt.Printf("%T, %p, %p\n", m, m, &m)

	m = make(map[string]int)
	fmt.Printf("%T, %p, %p\n", m, m, &m)
	m["name"] = 100

	fmt.Println(m)

	m1 := make(map[string]string)
	fmt.Printf("%T, %p, %p\n", m1, m1, &m1)

	m1["name"] = "Tom"
	fmt.Println(m1)
}
