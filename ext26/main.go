package main

import (
	"fmt"
)

//利用递归方法求 5!。 递归公式：fn = fn(1)*4!。
func fact(n int) (sum int) {
	if n == 0 {
		return 1
	}

	sum = n * fact(n-1)
	return
}
func main() {
	fmt.Printf("5的阶乘为：%d\n", fact(5)) //5的阶乘为：120
}
