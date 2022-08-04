package main

import "fmt"

// 将一个正整数分解质因数。例如：输入 90，打印出 90 = 2 * 3 * 3 * 5。
func main() {
	var n, i int = 0, 0
	fmt.Printf("请输入一个数:")
	fmt.Scanf("%d\n", &n)
	fmt.Printf("%d = ", n)
	for i = 2; i <= n; i++ {
		for n != i {
			if n % i == 0 {
				fmt.Printf("%d * ", i)
				n = n / i
			} else {
				break
			}
		}
	}
	fmt.Printf("%d\n", n)
}

