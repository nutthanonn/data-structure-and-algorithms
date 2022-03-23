// fibonacci Problem

package main

import "fmt"

var ff = make(map[int]int)

func fib(n int) int {
	res := 0
	if val, ok := ff[n]; ok { //dynamic programming
		return val
	}
	if n <= 2 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	ff[n] = res
	return ff[n]
}

func main() {
	fmt.Println(fib(10))
}
