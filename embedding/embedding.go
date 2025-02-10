package embedding

import (
	"fmt"
)

type Base struct {
	Name string
	Age  int
}

type EmbedBase struct {
	Base
	Price int
}

type EmbedPointerBase struct {
	*Base
	Price int
}

func (b Base) printBase() {
	fmt.Println(b.Name, b.Age)
}

func changeBase(baseinfo EmbedPointerBase) {
	baseinfo.Name = "changed name"
	baseinfo.Age = 3
}

func Embedding() {
	baseinfo := Base{"name", 2}
	embedbaseinfo := EmbedPointerBase{&baseinfo, 30000}
	embedbaseinfo.printBase()

	changeBase(embedbaseinfo)
	baseinfo.printBase()
}