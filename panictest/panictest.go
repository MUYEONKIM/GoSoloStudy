package panictest

import (
	"fmt"
)

func CubeRoot(x int) {
	ctn := 1
	for i := 0; i < x; i++ {
		ctn++
		fmt.Println(ctn)

		if ctn%10 == 0 {
			panic("종료")
		}
	}


}

func Panictest() {
	CubeRoot(1000)
}
