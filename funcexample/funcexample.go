package funcexample

import (
	"fmt"
)

func modifyNum(arg []int) {
	for i := range arg {
		arg[i] *= 2
	}
}

// 함수를 쓰는데 필요한 예제입니다.
func Funcexample() {
	test := []int{1,2,3,4,5}
	fmt.Println(test, "첫번째")
	
	modifyNum(test)

	fmt.Println(test, "두번째")

}

func NamedTest() (value, nextPos int) {
	return value, nextPos
}