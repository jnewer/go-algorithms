package main

import (
	"fmt"
)

// 输出国际象棋棋盘。
func main() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf("■")
			} else {
				fmt.Printf("□")
			}
		}
		fmt.Printf("\n")
	}
}
