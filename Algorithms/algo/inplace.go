package main

import (
	"fmt"
	"math"
)

func reverseArray(s []byte) {
	n := len(s)
	for i := 0; i < int(math.Floor(float64(len(s)/2))); i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
}

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseArray(s)
	fmt.Println(string(s))
}
