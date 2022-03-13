package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func (h *MaxHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.maxHeapifly(len(h.arr) - 1)
}

func (h *MaxHeap) maxHeapifly(index int) {
	for h.arr[parent(index)] < h.arr[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapiflyDown(index int) {
	lastIndex := len(h.arr) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.arr[l] > h.arr[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.arr[index] < h.arr[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}

	}

}

func (h *MaxHeap) Extract() int {
	extracted := h.arr[0]
	l := len(h.arr) - 1
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	if len(h.arr) == 0 {
		fmt.Println("Error lenght 0")
		return -1
	}
	h.maxHeapiflyDown(0)

	return extracted
}

// get parent Node
func parent(i int) int {
	return (i - 1) / 2
}

// get left Node
func left(i int) int {
	return 2*i + 1
}

// get right Node
func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	fmt.Println("-----------------------------")
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)

	}

}
