package main

import (
	"fmt"
	"sync"
)

type ConcCounter struct {
	counter int
	sync.RWMutex
}

func (cc *ConcCounter) Increment() {
	cc.Lock()         //блокируем монитор
	defer cc.Unlock() //разблокировка
	cc.counter++      //инкремент
}

func main() {
	var wg sync.WaitGroup         //создаем вэйтгруп для ожидания завершения всех горутн
	concCounter := &ConcCounter{} //создаем конкурентный счетчик

	for i := 0; i < 6000; i++ { //запускаем 6000 горутин, инкрементирующих счетчик
		wg.Add(1)
		go func() {
			defer wg.Done()

			concCounter.Increment()
		}()
	}
	wg.Wait() //Ожидаем выполнения
	fmt.Println(concCounter.counter)
}
