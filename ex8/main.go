package main

import "fmt"

// 输出 9*9 乘法口诀表。
func main() {
	for i := 0; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("%d*%d=%-3d", j, i, i*j) /*-3d 表示左对齐，占 3 位*/
		}
		fmt.Printf("\n")
	}
}
