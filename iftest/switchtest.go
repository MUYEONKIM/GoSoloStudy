package iftest

import (
	"fmt"
	"reflect"
)

func CheckType(arg any) any {
	return reflect.TypeOf(arg)
}

func Switchtest() {
	var test interface{} = CheckType('q')

	switch test.(type) {
		case int: 
			fmt.Println("int")
		case string: 
			fmt.Println("string")
		default: 
			fmt.Println("default")
	}

}