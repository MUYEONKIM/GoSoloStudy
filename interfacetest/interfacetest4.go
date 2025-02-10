package interfacetest

import (
	"fmt"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Horn")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("beepbo")
}

func (r Robot) Walk() {
	fmt.Println("powering legs")
}

func Interfacetest4() {
	play(Whistle(""))
	play(Horn(""))

	var robotTest NoiseMaker = Robot("")

	robotTest.MakeSound()
	// 아래는 interface에 정의가 되어 있지 않아 호출이 불가능해진다
	// robotTest.Walk()
}