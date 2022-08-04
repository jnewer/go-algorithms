package main

import (
	"fmt"
	"math"
)

// 一个 5 位数，判断它是不是回文数。即 12321 是回文数，个位与万位相同，十位与千位相同。
func main() {
	var x, i int
	var max = 5
	fmt.Printf("请输入一个数：")
	fmt.Scanf("%d\n", &x)
	for i = 0; i < max/2; i++ {
		h := x % int(math.Pow10(max-i)) / int(math.Pow10(max-i-1))
		l := x % int(math.Pow10(i+1)) / int(math.Pow10(i))
		if h != l {
			fmt.Printf("%d 不是一个回文数.\n", x)
			break
		}
	}
	if i == max/2 {
		fmt.Printf("%d 是一个回文数.\n", x)
	}

}
