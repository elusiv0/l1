package main

import "fmt"

func main() {
	set := make(map[string]bool)
	objects := []string{"dog", "dog", "cat", "cat", "cat", "tree", "shark", "bird", "bird"}

	for _, val := range objects {
		set[val] = true
	}

	for k := range set {
		fmt.Printf("%s ", k)
	}
}
