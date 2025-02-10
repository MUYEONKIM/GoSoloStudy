package channeltest

import "fmt"

func Channeltest() {

	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	test := <-done

	fmt.Println(test, done)

	ch := make(chan string)

	go func() {
		ch <- "wqe"
	}()

	var i string
	i = <-ch
	fmt.Println(i)
}