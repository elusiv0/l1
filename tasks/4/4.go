package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	workers := flag.Int("workers", 10, "count of workers") //получаем флаг количества воркеров
	flag.Parse()
	dataChannel := make(chan int) //создаем каннал данных

	var wg sync.WaitGroup                                       //объявляем вэйтгруппу для последующего ожидания корректного завершения задач воркеров
	signalChannel := make(chan os.Signal, 1)                    //объявляем канал для получения сигнала прерывания
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM) //регистрируем канал по этому сигналу

	//запускаем воркеров
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go worker(dataChannel, &wg)
		fmt.Printf("%d worker has started\n", i+1)
	}

	i := 0
	for {
		select {
		case <-signalChannel: //получаем сигнал о завершении
			fmt.Println("Closing data channel")
			close(dataChannel) //закрываем канал, давая сигнал воркерам о том, что больше ничего не придет
			fmt.Println("Wait until workers do their tasks")
			wg.Wait() //ждем дозавершения их текущих задач
			fmt.Println("Workers has stopped")
			return //останавливаем выполнение
		default:
			fmt.Printf("Sending %d to channel\n", i)
			dataChannel <- i //отправляем в канал данные
			i++
		}
	}
}

func worker(sc <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for elem := range sc {
		fmt.Printf("Receive %v from channel\n", elem) //выыводим полученное значение
	}
}
