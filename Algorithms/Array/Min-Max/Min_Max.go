// divide and conquer

package main

import "fmt"

func MaxMin(low, high int, arr []int) (int, int) {
	arr_max := arr[low]
	arr_min := arr[low]

	if low == high { // if size of array == 1
		arr_max = arr[low]
		arr_min = arr[low]
		return arr_max, arr_min

	} else if high == low+1 { // if size of array == 2

		if arr[low] > arr[high] {
			arr_max = arr[low]
			arr_min = arr[high]
		} else {
			arr_max = arr[high]
			arr_min = arr[low]
		}
		return arr_max, arr_min

	} else {
		mid := int((low + high) / 2)
		arr_max1, arr_min1 := MaxMin(low, mid, arr)
		arr_max2, arr_min2 := MaxMin(mid+1, high, arr)

		if arr_max1 < arr_max2 {
			arr_max1 = arr_max2
		}

		if arr_min1 > arr_min2 {
			arr_min1 = arr_min2
		}

		return arr_max1, arr_min1
	}
}

func main() {
	number := []int{1000, 11, 445, 1, 330, 3000}
	max, min := MaxMin(0, len(number)-1, number)
	fmt.Println(max, min)
}
