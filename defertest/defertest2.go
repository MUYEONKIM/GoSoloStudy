package defertest

import "fmt"

func Defertest2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}