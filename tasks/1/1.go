package main

import (
	"fmt"
)

// структура Human с методами
type Human struct {
	name      string
	surname   string
	height    int
	weight    int
	walkSpeed int
	language  string
}

func (h *Human) GetSpeed() int {
	return h.walkSpeed
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetLanguage() string {
	return h.language
}

// структура Action с встроенной Human
type Action struct {
	Human
}

func (a *Action) Walk() {
	fmt.Printf("%s is walking with speed %d km/h\n", a.GetName(), a.GetSpeed())

}

func (a *Action) Talk() {
	fmt.Printf("%s is talking on %s\n", a.GetName(), a.GetLanguage())
}

func main() {
	act := &Action{
		//встраивание структуры Human
		Human{
			name:      "John",
			surname:   "Black",
			height:    176,
			weight:    65,
			walkSpeed: 7,
			language:  "Chinese",
		},
	}
	//проверка встроенной структуры вызовом методов этой структуры
	fmt.Printf("Human with name %s, speed %d, language %s", act.GetName(), act.GetSpeed(), act.GetLanguage())

	//вызов методов структуры Action
	act.Talk()
	act.Walk()
}
