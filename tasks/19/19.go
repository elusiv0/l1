package main

import "fmt"

func main() {
	in := "главрыба"
	runes := []rune(in)
	l := 0
	r := len(runes) - 1

	for r > l {
		runes[l], runes[r] = runes[r], runes[l]
		r--
		l++
	}

	fmt.Println(string(runes))
}
