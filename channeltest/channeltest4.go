package channeltest

import "fmt"

func Channeltest4() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2

	close(ch)
	
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// ch <- 3 'send on closed channel' 에러 발생
	
	a,b := <- ch

	println(a,b, "Qwe")
}