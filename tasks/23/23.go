package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	i := 2
	up := slice[:i]     //записываем в временную переменную слайс ограниченный сверху
	down := slice[i+1:] //записываем слайс ограниченный снизу

	modified := append(up, down...) //складываем два слайса в один и получаем слайс с удаленной переменной

	fmt.Println(modified)
}
