package grammar

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	switchMethod(2)
}

func switchMethod(a int) {
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("else")
	}
}

func typeSwitch() {
	var x interface{}
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("else")
	}
}
