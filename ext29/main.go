package main

import (
	"fmt"
	"math"
)

func out(n, x int) {
	if n > 1 {
		out(n-1, x)
	}

	r := x % int(math.Pow10(n)) / int(math.Pow10(n-1))
	fmt.Printf("%d", r)
}

//给一个不多于 5 位的正整数，要求：一、求它是几位数，二、逆序打印出各位数字。
func main() {
	var x int
	fmt.Printf("请输入一个数：")
	fmt.Scanf("%d\n", &x)
	for i := 5; i > 0; i-- {
		r := x / int(math.Pow10(i-1))
		if r > 0 {
			fmt.Printf("%d是一个%d位数字.\n", x, i)
			out(i, x)
			fmt.Printf("\n")
			break
		}
	}

}
