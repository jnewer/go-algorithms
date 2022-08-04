package main

import (
	"fmt"
)

// 有一分数序列：2/1，3/2，5/3，8/5，13/8，21/13…求出这个数列的前 20 项之和。
func main() {
	number := 20
	a, b, s := 2.0, 1.0, 0.0
	for n := 1; n <= number; n++ {
		s = s + a/b
		a, b = a+b, a
	}
	fmt.Printf("sum is %9.6f\n", s) //sum is 32.660261
}
