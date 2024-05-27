package main

import "fmt"

func main() {
	first := []int{1, 7, 5, 9, 19, 8}
	second := []int{1, 8, 89, 2, 3, 19, 99}
	var answ []int

	existingElems := make(map[int]bool) //мапа для хранения встреченных значений

	for _, val := range first {
		existingElems[val] = true
	}

	for _, val := range second {
		if _, ok := existingElems[val]; ok { //проверка находимости в сете этого значения, если есть - добавляем в слайс ответа
			answ = append(answ, val)
		}
	}
	fmt.Printf("Первый массив: %v\n", first)
	fmt.Printf("Второй массив: %v\n", second)
	fmt.Printf("Пересечение элементов: %v", answ)
}
