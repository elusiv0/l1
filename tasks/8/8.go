package main

import "fmt"

func main() {
	var digit int64 = 257

	fmt.Printf("Число до изменения бита: %d\n", digit)

	digit = changeBit(digit, 0, 0)

	fmt.Printf("Число после изменения бита: %d\n", digit)
}

func changeBit(d int64, position int, value int) int64 {
	mask := int64(1 << position) //сдвигаем единицу на {position} позиций
	switch value {
	case 1:
		d |= mask //если единица - используем побитовое OR
	case 0:
		mask = ^mask //инвертируем маску
		d = d & mask //побитовое AND
	}

	return d
}
