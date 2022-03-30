package main

import "fmt"

func binarySearch(arr []int, target int) int {
	var first = 0
	var last = len(arr) - 1
	for first <= last {
		var midPointer = int((first + last) / 2)
		var midValue = arr[midPointer]

		if midValue == target {
			return midPointer
		} else if midValue < target {
			first = midPointer + 1
		} else {
			last = midPointer - 1
		}
	}

	return -1
}

func main() {
	n := []int{2, 9, 11, 21, 22, 32, 36, 48, 76}

	fmt.Println(binarySearch(n, 32))
}
