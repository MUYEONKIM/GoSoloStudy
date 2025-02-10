package channeltest

import "fmt"

func Channeltest5() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 방법 1
	// 채널이 닫힌 것을 감지할 때 까지 계속 수신
	for {
		if i, success := <-ch; success {
			fmt.Println(i)
		} else {
			break
		}
	}

	// 방법2
	// 채널 range 문
	for i := range ch {
		fmt.Println(i)
	}

}