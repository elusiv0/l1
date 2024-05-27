package main

import (
	"fmt"
	"time"
)

func main() {
	sleep(2 * time.Second)
	fmt.Println("I am awake")
}

func sleep(dur time.Duration) {
	t := time.NewTimer(dur) //создаем новый таймер с указанным временем
	<-t.C                   //ожидаем завершения таймера
	fmt.Println("timer fired")
}
