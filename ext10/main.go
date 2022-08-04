package main

import (
	"fmt"
)

// 打印楼梯，同时在楼梯上方打印笑脸。
func main() {

	for i := 0; i < 8; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("■■")
		}
		fmt.Printf("☺\n")
	}
}
