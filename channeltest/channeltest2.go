package channeltest

import "fmt"

func Channeltest2()  {
	// c := make(chan int)

	// // 수신 루틴이 없으므로 데드락 에러가 발생한다.
	// c <- 1

	// // 코멘트해도 데드락이 걸린다 (별도의 go 루틴이 없기 때문)
	// fmt.Println(<-c) 

	// 아래와 같이 해도 데드락이 걸림, 버퍼크기는 9인데 값을 10번 보내기 때문
	c2 := make(chan int, 9)


	for i := 0 ; i < 10; i++ {
		c2 <- i
	}

	fmt.Println(<-c2)
}