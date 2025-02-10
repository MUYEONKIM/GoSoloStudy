package interfaceerror

import "fmt"

type overHeat float64

func (o overHeat) Error() string {
	return fmt.Sprintf("over heat is : %0.2f", float64(o))
}

func checkTemp(act float64, crit float64) error {
	axcess := act - crit

	if axcess > 0 {
		temp := overHeat(axcess)
		return temp
	}
	return nil
}

func Interfaceerror() {
	err := checkTemp(38.5, 37.5)

	if err != nil {
		fmt.Println(err)
	}
}