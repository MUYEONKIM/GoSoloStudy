package pass

import "fmt"

type Person struct {
	name string
	age int
	gender bool
}

func changeName(p Person) Person {
	p.name = "구마적"
	return p
}

func Pass() {
	p := Person{name: "김두한", age : 30, gender : true}
	copy := changeName(p)
	fmt.Printf("복사 = %v\n", copy)
	fmt.Printf("원본 = %+v\n", p)
}