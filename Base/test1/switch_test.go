package test

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	checkNum(5)
	checkNum2(6)
	checkNum3(11)
}

func checkNum(num int) {
	switch num {
	case 1:
		fmt.Println("等于1")
	case 2:
		fmt.Println("等于2")
	case 3:
		fmt.Println("等于3")
	default:
		fmt.Println("其它值", num)
	}
}

func checkNum2(num int) {
	switch num {
	case 1, 5, 10:
		fmt.Println("等于1, 5, 10")
	case 2, 6:
		fmt.Println("等于2, 6")
	case 3:
		fmt.Println("等于3")
	default:
		fmt.Println("其它值", num)
	}
}

func checkNum3(num int) {
	switch {
	case num >= 1 && num <= 5:
		fmt.Println("1到5之间")
	case num >= 6 && num <= 10:
		fmt.Println("6到10之间")
	case num > 10:
		fmt.Println("大于10")
	default:
		fmt.Println("其它值", num)
	}
}
