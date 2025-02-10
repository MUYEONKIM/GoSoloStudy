package datatest

import (
	"fmt"
)

type Ttype struct {
	arg1 int
	arg2 bool
}

func Datatest() {
	result := new(Ttype)
	var result2 Ttype 

	fmt.Println(result, "1")
	fmt.Println(result2, "2")

	test1 := new(Ttype)
	test2 := &Ttype{}

	fmt.Println(test1, test2)

	qq := make([]int, 10, 100)
	fmt.Println(qq, "이거지")

	qw := [...]int{1,2,3,4,5,6}

	fmt.Printf("%T\n", qw)




}
