package gorutine

import (
	"fmt"
	"sync"
	"time"
)

func Say(arg string) {
	fmt.Println(arg)
}

func Gorutin() {
	Say("1번")

	go Say("2번")
	go Say("3번")
	go Say("4번")
	go Say("5번")

	time.Sleep(1 * time.Second)
}

func Goroutine() {
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		fmt.Println("Done")
	}()

	go func(msg string) {
		defer wait.Done()
		fmt.Println(msg)
	}("Hi")

	wait.Wait()
}