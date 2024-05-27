package main

import "fmt"

func main() {
	a := 99
	b := 198
	fmt.Printf("Числа до свапа: a = %d b =% d\n", a, b)
	a = a - b
	b = a + b
	a = b - a
	fmt.Printf("Числа после свапа: a = %d, b = %d", a, b)
}
