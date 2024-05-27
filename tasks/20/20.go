package main

import (
	"fmt"
	"strings"
)

func main() {
	in := "snow dog sun — sun dog snow"

	w := strings.Split(in, " ")
	l := 0
	r := len(w) - 1

	for r > l {
		w[l], w[r] = w[r], w[l]
		r--
		l++
	}

	fmt.Println(strings.Join(w, " "))
}
