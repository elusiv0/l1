package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

// При ограничении крайнего элемента для получения новой строки из изначальной, ссылка на байты изначальной будут хранится и далее, что нехорошо с точки зрения расхода памяти
// Следовательно, лучше выделить место под новую строку с заданной размерностью, чтобы Go мог затереть данные изначальной большой строки
// Для этого можно воспользоваться методом string.Clone()
var justString string

func main() {
	var m1, m2 runtime.MemStats
	setStr()
	runtime.ReadMemStats(&m1)
	fmt.Println(m1.Alloc)
	runtime.GC()
	time.Sleep(2 * time.Second)
	runtime.ReadMemStats(&m2)
	fmt.Println(m2.Alloc)
}

func setStr() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func createHugeString(sz int) string {
	var b strings.Builder
	b.Grow(sz)
	for i := 0; i < sz; i++ {
		b.WriteByte(byte(97 + rand.Intn(23)))
	}

	return b.String()
}
