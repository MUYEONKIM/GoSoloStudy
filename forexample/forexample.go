package forexample

import "fmt"

func Forexample() {
	sum := 0

	for x, y := 1, 3; x < 10; x , y = x+2, y+2 {
		sum += x
		sum -= y
	}

	// L1:
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// 	if i == 5 {
	// 		break L1
	// 	}	
	// }

	fmt.Println("합계 = " ,sum)
    // var a = 1
    // for a < 15 {
    //     if a == 5 {
    //         a += a
    //         continue // for루프 시작으로
    //     }
    //     a++
    //     if a > 10 {
    //         break  //루프 빠져나옴
    //     }
    // }
    // if a == 11 {
    //     goto END //goto 사용예
    // }
    // fmt.Println(a)
 
	// END:
    // fmt.Println("End")

	// name := []string{"김두한", "구마적", "쌍칼"}

	// for i, e := range name {
	// 	fmt.Println(i, " ",e)
	// }

	// i := 1

	// for {
	// 	sum += i
	// 	i++

	// 	if i == 100 {
	// 		break
	// 	}
	// }

	// for i < 101 {
	// 	sum += i
	// 	i++
	// }
	// for i := 1; i < 11; i++ {
	// 	sum += i
	// }

	// fmt.Println(sum, "합계")
}