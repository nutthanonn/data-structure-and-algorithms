/*
	Kadane's Algorithm
*/

package main

import "fmt"

func Kadane(arr []int) int {
	max_curr, max_global := arr[0], arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] > max_curr+arr[i] {
			max_curr = arr[i]
		} else {
			max_curr += arr[i]
		}
	}

	if max_curr > max_global {
		max_global = max_curr
	}
	return max_global
}

func main() {
	arr := []int{-2, 3, 2, -1}
	fmt.Println(Kadane(arr))
}
