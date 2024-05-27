package main

import (
	"sync"
)

type ConcMap struct {
	mymap map[string]int
	sync.RWMutex
}

func New() *ConcMap {
	concMap := ConcMap{
		mymap: make(map[string]int),
	}

	return &concMap
}

func (concMap *ConcMap) Set(key string, val int) {
	concMap.Lock()         //блокируем объект
	defer concMap.Unlock() //разблокировка
	concMap.mymap[key] = val
}

func (concMap *ConcMap) Get(key string) (int, bool) {
	concMap.RLock()         //блокируем объект на все .Lock(), пропуская .RLock() для параллельного чтения
	defer concMap.RUnlock() //разблокировка
	val, ok := concMap.mymap[key]

	return val, ok
}
