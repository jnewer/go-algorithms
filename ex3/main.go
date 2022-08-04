package main

import (
	"fmt"
	"math"
)

//一个整数，它加上 100 后是一个完全平方数，再加上 268 又是一个完全平方数，请问该数是多少？
func main() {
	i := 0
	for {
		x := int(math.Sqrt(float64(i + 100)))
		y := int(math.Sqrt(float64(i + 268)))
		if x*x == i+100 && y*y == i+268 {
			fmt.Printf("这个数字是 %d\n", i)
			break
		}
		i++
	}
}
