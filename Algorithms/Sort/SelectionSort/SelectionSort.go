package main

import "fmt"

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	n := []int{2, 31, 32, 12, 4553, 123, 56, 2, 1, 354, 87, 1}
	fmt.Println(SelectionSort(n))
}
