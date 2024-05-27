package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 4, 8, 16, 32, 64} //исходный массив
	dataChannel := make(chan int)         //канал для записи значений из массива
	doubleChannel := make(chan int)       //канал для записи удвоенных значений

	//горутина для получения значения из канала, удвоения и записи в следующий канал
	go func() {
		for val := range dataChannel {
			doubleChannel <- val * 2 //удваиваем значений и записываем в канал
		}
		close(doubleChannel) //закрываем канал после записи всех значений
	}()

	//горутина для записи в канал значений для последующей обработки
	go func() {
		for _, val := range data {
			dataChannel <- val
		}
		close(dataChannel) //закрываем канал, сигнализируя о том, что ничего больше не придет и ожидать более не следует
	}()

	for val := range doubleChannel {
		fmt.Printf("%d ", val)
	}
}
