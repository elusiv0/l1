package main

import "fmt"

//интерфейс адаптера
type MoveAdapter interface {
	Move()
}

//конкретные реализации, к которым мы хотим обратиться(Human и Bird)
type Human struct{}

func (h *Human) Walk() {
	fmt.Println("Human is moving by walking")
}

type Bird struct{}

func (b *Bird) Fly() {
	fmt.Println("Bird is moving by flying")
}

//Адаптеры для реализаций, реализующие контракт интерфейса MoveAdapter
type HumanAdapter struct {
	*Human
}

func NewHumanAdapter(h *Human) MoveAdapter {
	return &HumanAdapter{
		h,
	}
}

func (ha *HumanAdapter) Move() {
	ha.Walk()
}

type BirdAdapter struct {
	*Bird
}

func NewBirdAdapter(b *Bird) MoveAdapter {
	return &BirdAdapter{
		b,
	}
}

func (b *BirdAdapter) Move() {
	b.Fly()
}

func main() {
	human := &Human{}
	hAdapter := NewHumanAdapter(human)
	bird := &Bird{}
	bAdapter := NewBirdAdapter(bird)

	example(hAdapter)
	example(bAdapter)
}

//пример полиморфного вызова конечной реализации движения
func example(adapter MoveAdapter) {
	adapter.Move()
}
