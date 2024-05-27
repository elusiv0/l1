package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	stopCh := make(chan bool, 1)
	var wg sync.WaitGroup

	//остановка горутины с помощью канала сигнала об остановке
	wg.Add(1)
	go func(stopChannel chan bool) {
		defer wg.Done()

		for {
			select {
			case <-stopChannel:
				fmt.Println("get signal to stop goroutine from channel...")
				return
			default:
				fmt.Println("I am working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(stopCh)

	time.Sleep(1 * time.Second)
	stopCh <- true //сигнализируем горутине об остановке посредством записи в канал
	wg.Wait()

	//остановка горутины при помощи передачи в функцию контекста с имеющейся функицей cancel у этого контекста, полученной посредством context.WithCancel()
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done(): //получаем сигнал о завершении контекста
				fmt.Println("get signal to stop goroutine from context....")
				return
			default:
				fmt.Println("I am working")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancel() //вызываем функции с сигналом о завершении
	wg.Wait()

	channel := make(chan int)

	//остановка горутины посредством закрытия канала
	wg.Add(1)
	go func(channel chan int) {
		defer wg.Done()

		for {
			select {
			case <-channel:
				fmt.Println("Get signal from closing channel, goroutine will be stopped...")
				return
			default:
				fmt.Println("I am working...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}(channel)

	time.Sleep(1 * time.Second)
	close(channel) //закрываем канал для сигнализации об остановке
	wg.Wait()

	//остановка горутины посредством передачи в нее флага и проверки на выставление этого флага на завершении горутины
	var stopFlag bool
	wg.Add(1)
	go func(flag *bool) {
		defer wg.Done()

		for {
			if *flag {
				fmt.Println("Get signal from flag, goroutine will be stopped...")
				return
			}
			fmt.Println("I am working...")
			time.Sleep(200 * time.Millisecond)
		}
	}(&stopFlag)

	time.Sleep(1 * time.Second)
	stopFlag = true //сигнализируем горутине посредтвом флага
	wg.Wait()
}
