package main

import "fmt"

func BubbleSort(arr []int, s int) []int {
	if s == 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	BubbleSort(arr, s-1) // recursive function
	return arr
}

func main() {
	arr := []int{2, 1, 4, 13, 2, 90, 15, 4}
	fmt.Println(BubbleSort(arr, len(arr)))
}
