package interfacetest

import "fmt"

type myInterface interface {
    methodWithoutParameters()
    methodWithParameter(float64)
    methodWithReturnValue() string
}

type myType int

func (m myType) methodWithoutParameters() {
	fmt.Println("mwithoutparmas")
}

func (m myType) methodWithParameter(arg float64) {
	fmt.Println("mwparmas", arg)
}

func (m myType) methodWithReturnValue() string {
	fmt.Println("mwrvalue")
	return "end"
}

func Interfacetest2() {
	// myinterface에 값을 ekadk wnftneh dlTdma
	var value myInterface = myType(15)

	value.methodWithParameter(25)

	fmt.Println(value, "Qwer")

}