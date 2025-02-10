package anonymoustype

import "fmt"

func Namedtyped() {
	type Namedtype struct {
		count int
	}

	var test2 struct {
		count int
	}

	var test1 Namedtype

	test1 = test2

	fmt.Printf("%v\n", test1)
}