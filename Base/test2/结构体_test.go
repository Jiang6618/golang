package test

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	type student struct {
		Name  string
		Age   int
		Score [5]float32
		Ptr   *int
		slice []int
		map1  map[string]string
	}

	var s student
	fmt.Printf("%#v", s)

	s.slice = make([]int, 10)
	s.map1 = make(map[string]string, 5)
	s.map1["key1"] = "value"

	fmt.Printf("%#v", s)
}
