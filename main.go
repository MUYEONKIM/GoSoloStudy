package main

import (
	// "GOSTUDY/anonymoustype"
	// "GOSTUDY/datatest"
	// "GOSTUDY/defertest"
	// "GOSTUDY/forexample"
	// "GOSTUDY/funcexample"
	// "GOSTUDY/maketest"
	// "GOSTUDY/pass"
	// "GOSTUDY/testtest"

	// "GOSTUDY/embedding"
	// "GOSTUDY/errorhandle"
	"GOSTUDY/gorutine"
	"runtime"

	// "GOSTUDY/interfaceerror"
	// "GOSTUDY/interfacefinal"
	// "GOSTUDY/interfacestringer"

	// "GOSTUDY/panictest"

	"fmt"
	// "unsafe"
)

// 다른 패키지에서도 쓰려면 앞글자 대문자
type person struct {
	name string
	age int
	gender bool
}

func main() {
	// // 객체 생성
	// p := person{}

	// // 필드값 설정
	// p.name = "김두한"
	// p.age = 10
	// p.gender = true

	// // 초기화와 동시에 선언해주어도 된다, 
	// // struct의 값을 지정안해줄 경우 zero value
	// // 항상 구조체는 선언되어 있어야 할 것
	// p2 := person{
	// 	name : "구마적",
	// 	age : 25,
	// }

	// // 내장함수 new()를 이용해서 선언, 모든 필드 zero value
	// p3 := new(person)

	// fmt.Printf("%+v\n",p)
	// // {name:김두한 age:0 gender:true}
	// fmt.Printf("%+v\n", p2)
	// // {name:구마적 age:25 gender:false}
	// fmt.Printf("%+v\n", p3)
	// // &{name: age:0 gender:false}, 현재 형식이 %v라서 적절한 값으로 출력
	// print(p3)
	// // 0xc00005e3c0 위와 같이 할 경우 포인터이기 때문에 주소값이 출력
	// // p := Person{name: "김두한", age : 30, gender : true}

	// fmt.Printf("x = %p\n", p3)

	// pass.Pass()
	// pass.Passbyref()
	// anonymoustype.Anonymoustype()
	// anonymoustype.Namedtyped()

	// type qq struct {
	// 	string
	// 	int32
	// }

	// fmt.Println(unsafe.Sizeof(qq{}))

	// forexample.Forexample()

	// funcexample.Funcexample()
	// value, nextPos := funcexample.NamedTest()
	// fmt.Println(value, nextPos, "QWE")

	// defertest.Defertest()
	// defertest.Defertest2()
	// defertest.Defertest3()

	// datatest.Datatest()

	// result := maketest.Maketest()

	// testtest.TT()
	// fmt.Printf("%+v\n",result)
	// interfacetest.Interfacetest()

	// functiontest.Functiontest()
	// functiontest.Functiontest2()

	// methodtest.Methodtest()

	// interfacetest.Interfacetest2()
	// interfacetest.Interfacetest3()
	// interfacetest.Interfacetest4()

	// interfacefinal.Interfacefinal()
	// interfaceerror.Interfaceerror()

	// interfacestringer.Interfacestringer()

	// errorhandle.Errorhandle()

	// embedding.Embedding()

	// panictest.Panictest()
	// recovertest.Recovertest()

	// gorutine.Gorutin()
	runtime.GOMAXPROCS(2)
	gorutine.Goroutine()

}


func init() {
	fmt.Println("init")
}
