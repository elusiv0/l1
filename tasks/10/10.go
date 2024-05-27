package main

import "fmt"

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 59.9999}

	tempGroups := make(map[int][]float64) //мапа для хранения групп по началу диапазона температуры и слайса значений, входящих в указанный диапазон

	for _, val := range temp {
		key := int(val/10) * 10                        //получаем ключ (начало интервала)
		tempGroups[key] = append(tempGroups[key], val) //добавляем в слайс по этому ключу значение
	}

	for k, v := range tempGroups {
		fmt.Printf("%d:{", k) //выводим начало интервала температурной группы
		div := ""             //разделитель для вывода значений
		for _, val := range v {
			fmt.Printf("%s%v", div, val)
			div = ", "
		}
		fmt.Print("}\n")
	}
}
