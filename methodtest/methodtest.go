package methodtest

import "fmt"

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

func (r *Rect) areaPointer() int {
	r.width++
	return r.width * r.height
}

func Methodtest() {
	rect := Rect{10, 20}
	area := rect.area()
	fmt.Println(area)

	areaPointer := rect.areaPointer()
	fmt.Println(area, rect, areaPointer)

}