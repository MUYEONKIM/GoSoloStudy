package testtest

import "fmt"

func sums(arg ...int) int {
	fmt.Println("여기가 arg : ", arg)
	res := 0
	for _, n := range arg {
		res += n
	}
	return res
}

func TT() {
	testarg := []int{1,2,3}

	result := sums(testarg...)

	fmt.Println(result)
}
