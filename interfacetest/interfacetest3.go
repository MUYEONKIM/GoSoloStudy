package interfacetest

import "fmt"

type myInterface2 interface {
    methodWithoutParameters2()
    methodWithParameter2(float64)
    methodWithReturnValue2() string
}

type myType2 struct {}

func (m myType2) methodWithoutParameters2() {
	fmt.Println("mwithoutparmas")
}

func (m myType2) methodWithParameter2(arg float64) {
	fmt.Println("mwparmas", arg)
}

func (m myType2) methodWithReturnValue2() string {
	return "end"
} 

func Interfacetest3() {
	var temp myInterface2 = myType2{}
	temp.methodWithParameter2(2.3)


	fmt.Println(temp, "Qwe")

}