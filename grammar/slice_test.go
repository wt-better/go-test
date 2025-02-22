package grammar

import "testing"

func TestSlice(t *testing.T) {
	//array := [...]int{1, 2, 3}
	//
	//var slice []int = array[1:3]
	//
	//numbers := append(slice, 1)
	//
	////slice := make([]int, 10)
	////var sliceMake []int = make([]int,10)
	////sliceMake := make([]int, 3)
	////sliceMake = append(sliceMake, 1)
	////for i := range sliceMake {
	////	fmt.Println(sliceMake[i])
	////}
	//
	////copy(slice, sliceMake)
	//
	//for i := range numbers {
	//	fmt.Println(numbers[i])
	//}

	months := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	summer := months[6:9]
	summer[0] = 13

	for i := range months {
		println(months[i])
	}

	//
	//println("cap: " + strconv.Itoa(cap(summer)))
	//println("len: " + strconv.Itoa(len(summer)))
	//
	//for i := range summer {
	//	println(summer[i])
	//}
}
