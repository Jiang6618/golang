package test

import (
	"fmt"
	"testing"
)

func TestIf(test *testing.T) {
	if getOK() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if size := getSize(); size > 100 {
		fmt.Println("大于100")
	} else if size <= 100 && size >= 0 {
		fmt.Println("小于100大于0")
	} else {
		fmt.Println("小于0")
	}
}

func getSize() int {
	return 90
}

func getOK() bool {
	return true
}
