package main

import "fmt"

//输入三个 整数 x，y，z，请把这三个数由小到大输出。
func main() {
	var x, y, z int = 0, 0, 0
	fmt.Scanf("%d%d%d", &x, &y, &z)
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	fmt.Printf("%d < %d < %d\n", x, y, z)
}
