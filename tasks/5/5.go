package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	secs := flag.Int("seconds", 5, "number of seconds to stop") //получаем количество секунд таймаута
	flag.Parse()
	fmt.Printf("Running program with %d seconds timeout\n", *secs)
	dataChannel := make(chan int) //инициализация канала
	var wg sync.WaitGroup         //вэйтгруппа для ожидания завершения задачи горутины

	wg.Add(1)
	go consumer(dataChannel, &wg) //запускаем горутину

	i := 0
	timeout := time.After(time.Duration(*secs) * time.Second) //инициализируем таймер
	for {
		select {
		case <-timeout: //получение сигнала о завершении по таймауту
			fmt.Println("Stopping program after timeout")
			close(dataChannel) //закрываем канал
			fmt.Println("Channel has closed")
			wg.Wait() //ждем завершение таски
			fmt.Println("Worker has ended his task")
			return
		default:
			fmt.Printf("Sending %v to channel\n", i)
			dataChannel <- i //отправляем в канал значение
			i++
		}
	}
}

func consumer(dataChannel <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() //сигнализируем о конце

	for val := range dataChannel {
		fmt.Printf("Received %v from channel\n", val)
	}
}
