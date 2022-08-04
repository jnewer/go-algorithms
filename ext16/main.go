package main

import "fmt"

//输入两个正整数 m 和 n，求其最大公约数和最小公倍数。
func main() {
	var m, n, r, x int
	fmt.Print("请输入两个数：")
	fmt.Scanf("%d%d", &m, &n)
	x = m * n
	for n != 0 {
		r = m % n
		m = n
		n = r
	}
	fmt.Printf("最大公约数：%d,最小公倍数：%d\n", m, x/m)
}
