package maketest

import (
	"fmt"
	"log"
)

func Sum(a *[3]float64) (sum float64) {
	for i, v := range *a {
		sum += v
		fmt.Println("I : ", i, " v : ", v)
	}
	return sum
}

func Maketest() int {
	// 아래는 안된다. 왜냐하면 new는 포인터를 반환하기 때문에 타입이 *[]int가 되어야 한다.
	// var p []int = new([]int)
	// var p *[]int = new([]int)
	// // 아래는 용량 11중에서 10개만 나타내는 것
	// var v []int = make([]int, 10, 11)

	// fmt.Println(p)
	// fmt.Println(v)

	// array := [...]float64{2.0, 3.2, 1.6}
	// Sum(&array)

	// a := []int{1,2,3,4,5}
	// b := []int{1,2,3}

	// c := append(a, 6)
	// fmt.Println(c, a)

	// fmt.Println(cap(a), cap(b))

	// a := map[string]int {
	// 	"q" : 3,
	// 	"w" : 2,
	// }

	// fmt.Println(a, a["q"], a["e"])
	test := map[string]int {
		"Q" : 3,
		"w" : 3,
	}

	fmt.Println(test)

	delete(test, "Q")
	fmt.Println(test)
	delete(test, "Q")
	fmt.Println(test)
	delete(test, "w")
	fmt.Println(test)

	x := []int{1,2,3}
	y := []int{4,5,6}

	x = append(x, y...)

	testargs := []int{1,2,3}
	result := sums(testargs...)

	fmt.Println("여기가 result : ", result)

	return offset("qwe")


}

func sums(arg ...int) int {
    fmt.Println("여기가 arg : ", arg)
    res := 0
    for _, n := range arg {
        res += n
    }
    return res
}

func offset(tz string) int {
	timeZone := map[string]int {
		"UTC":  0*60*60,
		"EST": -5*60*60,
		"CST": -6*60*60,
		"MST": -7*60*60,
		"PST": -8*60*60,		
	}

	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknwon time Zone", tz)

	return 0
}