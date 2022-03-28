package main

import "fmt"

func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}

	return -1
}

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(LinearSearch(number, 2))
}
