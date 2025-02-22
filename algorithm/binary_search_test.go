package algorithm

import (
	"fmt"
	"testing"
)

func binarySearch(array *[9]int, left int, right int, val int) int {
	if left > right {
		return -1
	}

	var middle = (left + right) / 2
	if (*array)[middle] > val {
		return binarySearch(array, left, middle-1, val)
	} else if (*array)[middle] < val {
		return binarySearch(array, middle+1, right, val)
	} else {
		return middle
	}
}

//func main() {
//	var array = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Print(binarySearch(&array, 0, 9, 0))
//}

func TestBinarySearch(t *testing.T) {
	var array = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Print(binarySearch(&array, 0, 9, 0))
}
