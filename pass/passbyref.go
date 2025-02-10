package pass

import "fmt"

type PersonRef struct {
	name string
	age int
	gender bool
}

func changedName(p *PersonRef) {
	p.name = "구마적"
}

func Passbyref() {
	pref := PersonRef{
		name : "김두한",
		age : 23,
		gender : false,
	}

	changedName(&pref)

	fmt.Printf("참조시원본 = %+v\n", pref)
}