package functiontest

import "fmt"

func Functiontest() {
	sum := func(n ...int) int {
		result := 0
		for _, arg := range n {
			result += arg
		}
		return result
	}

	test := []int{1, 5, 2, 5, 6, 7, 4, 32, 2, 5}

	fmt.Println(sum(test...))

}

func Functiontest2() {
	add := func(i int, j int) int {
		return i + j
	}

	// 익명함수 변수에 담아 전달
	r1 := calc(add, 10, 20)
	fmt.Println(r1)

	// 직접 파라미터에 함수 정의
	r2 := calc(func(x int, y int) int { return x - y}, 50, 30)
	fmt.Println(r2)
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a,b)

	return result
}

type calculator func(int, int) int

func calcemit(f calculator, a int, b int) int {
	result := f(a,b)

	return result
}