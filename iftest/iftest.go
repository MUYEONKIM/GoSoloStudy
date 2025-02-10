package iftest

import (
	"fmt"
	"io/ioutil"
	"os"
)

func If()  {

	// 파일 열기
	f, err := os.Open("C:/test.txt")
	if test := 1 ;err != nil {
		fmt.Println("파일 열기 오류:", err)
		fmt.Println(test)

		test := 2
		fmt.Println(test)

		return
	}
	defer f.Close() // 함수가 끝나면 파일 닫기

	// 파일의 내용을 읽기
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("파일 읽기 오류:", err)
		return
	}


	// 파일 내용 출력
	fmt.Println("파일 내용:", string(content))

}