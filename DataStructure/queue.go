package main

import "fmt"

type Queue []int

func (q Queue) Push(val int) Queue {
	return append([]int{val}, q...)
}

func (q Queue) Pop() Queue {
	l := len(q)
	if l == 0 {
		fmt.Println("No queue")
	}
	return q[:l-1]
}

func main() {
	q := Queue{}
	q = q.Push(1)
	q = q.Push(2)
	q = q.Push(3)

	q = q.Pop()
	fmt.Println(q)
}
