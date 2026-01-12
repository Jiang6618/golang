package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFmt(t *testing.T) {
	var name string = "Tom"

	var age int = 18

	isRoot := true

	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	fmt.Printf("isRoot: %v\n", isRoot)
	fmt.Println("age: ", age)
}

func TestT(t *testing.T) {
	n1 := 1
	n2 := 2
	n3 := 1

	fmt.Println(n1 == n3 || reflect.TypeOf(n2).Kind() == reflect.String)
}
