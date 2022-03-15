package main

import "fmt"

// O(n^2)
func InsertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{10, 2, 3, 5, 1, 48, 90, 12, 21, 35, 46, 99}
	fmt.Println(InsertionSort(arr))
}
