package main

import "fmt"

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, int) {
	l := len(s)
	if l == 0 {
		fmt.Println("Error len stack: empty stack")
	}
	return s[:l-1], s[l-1]
}

func main() {
	s := make(stack, 0)
	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)

	s, p := s.Pop()
	fmt.Println(p) // ตัวที่ถูก pop
	fmt.Println(s) // ตัวที่เหลือ

}
