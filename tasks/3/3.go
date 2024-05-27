package main

import (
	"fmt"
	"sync"
)

func main() {
	inSlice := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup //waitgroup для последующего ожидания отработки горутин

	sum := make(chan int, 1) //создаем буферизированный канал для сумм квадратов
	sum <- 0                 //добавляем изначальный 0
	for i := range inSlice {
		wg.Add(1)         //добавляем горутину в вэйтгрупу
		val := inSlice[i] //достаем переменную из массива для последующей передачи в горутину

		go func() {
			defer wg.Done() //сигнализируем о завершении горутины

			curSum := <-sum        //достаем нынешнее значению суммы
			curSquare := val * val //возводим переменную в квадрат

			sum <- curSum + curSquare //записываем сумму плюс квадрат значения в канал сумм
		}()
	}

	wg.Wait()          //ждем отработки горутин
	close(sum)         //закрываем канал
	fmt.Println(<-sum) //выводим значение суммы квадратов в stdout
}
