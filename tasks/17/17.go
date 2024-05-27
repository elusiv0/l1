package main

import "fmt"

func main() {
	arr1 := []int{1, 4, 6, 9, 10, 11, 12}
	fmt.Println(binarySearch(arr1, 0))
}

func binarySearch(arr []int, target int) int {
	l := -1
	r := len(arr)

	for r-1 > l {
		mid := l + (r-l)/2
		if arr[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}

	return l
}
