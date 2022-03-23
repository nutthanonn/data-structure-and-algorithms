// fibonacci Problem

package main

import "fmt"

var ff = make(map[int]int)

func fib(n int) int {
	res := 0
	if val, ok := ff[n]; ok { //dynamic programming --> memoization
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

func fib2(n int) int {
	ff := make(map[int]int)
	ff[1] = 1
	ff[2] = 1

	for i := 3; i < n+1; i++ { //dynamic programming --> memoization (Bottom up)
		ff[i] = ff[i-1] + ff[i-2]
	}

	return ff[n]
}

func main() {
	fmt.Println(fib(1000))
	fmt.Println(fib2(1000))
}
