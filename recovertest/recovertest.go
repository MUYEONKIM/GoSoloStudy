package recovertest

import (
	"fmt"
	"os"
)

func openFile(fn string) {
	defer func() {
		if r:= recover(); r != nil {
			fmt.Println("OPEN ERROR : ", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	fmt.Println(f.Name())

	defer f.Close()
}

func Recovertest() {
	// 여기서 오류가 발생했지만, recover에 의해 복구된 후
	openFile("/tessdfsdfsdft.txt")

	// 정상적으로 아래의 코드 진행
	fmt.Println("Done")

}