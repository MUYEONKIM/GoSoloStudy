package errorhandle

import (
	"fmt"
	"log"
	"os"
)

func Errorhandle() {
	f, err := os.Open("/test.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(f.Name())
}