package anonymoustype

import "fmt"

type test1 struct {
	a int
	b int
}

type test2 struct {
	a int
	b int
}

func Anonymoustype() {
	var result1 test1
	var result2 test2

	result1 = test1(result2)


	fmt.Printf("같냐? = %v\n",result1 )

}