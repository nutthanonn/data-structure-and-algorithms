package main

import (
	"fmt"
	"math"
	"strconv"
)

func IntPow(n int, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func karatsuba(num1, num2 int) int {
	if num1 < 10 || num2 < 10 {
		return num1 * num2
	}

	m := max(len(strconv.Itoa(num1)), len(strconv.Itoa(num2)))
	m2 := int(math.Ceil(float64(m / 2)))

	a := int(math.Ceil(float64(num1 / IntPow(10, m2))))
	b := int(math.Ceil(float64(num1 % IntPow(10, m2))))
	c := int(math.Ceil(float64(num2 / IntPow(10, m2))))
	d := int(math.Ceil(float64(num2 % IntPow(10, m2))))
	z0 := karatsuba(b, d)
	z1 := karatsuba((a + b), (c + d))
	z2 := karatsuba(a, c)

	return (z2 * IntPow(10, (2*m2))) + ((z1 - z2 - z0) * IntPow(10, (2*m2))) + (z0)
}

func main() {
	fmt.Println(karatsuba(4, 58779))
}
