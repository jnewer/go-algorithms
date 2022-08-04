package main

import (
	"fmt"
)

// 古典问题：有一对兔子，从出生后第 3 个月起每个月都生一对兔子，
// 小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？
func main() {
	var m1, m2 int = 1, 1
	for i := 1; i < 12; i++ {
		fmt.Println(m1, m2)
		m1 += m2
		m2 += m1
	}
}
