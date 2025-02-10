package defertest

import (
	"fmt"
	"os"
)

func Defertest() {
	file, err := os.Open("/test.txt")

	if err != nil {
		fmt.Println("파일 열기 오류 : ", err)
		return
	}

	fmt.Println(file)

	defer fmt.Println("언제 실행되게?")


	fmt.Println("완료")
	return 
}