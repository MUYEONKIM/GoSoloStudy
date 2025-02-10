package interfacestringer

import "fmt"

type coffee string

func (c coffee) String() string {
	return string(c) + "coffee pot"
}

func Interfacestringer() {
	coffeepot := coffee("cold brew")
	fmt.Println(coffeepot)
}