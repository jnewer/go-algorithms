package main

import (
	"fmt"
	"math"
)

//判断 101-200 之间有多少个素数，并输出所有素数。
func main() {
	var i, j, k, count int = 0, 0, 0, 0
	for i = 101; i < 200; i++ {
		k = int(math.Sqrt(float64(i)))
		for j = 2; j <= k; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == k+1 {
			fmt.Print(i, " ")
			count++
		}
	}

	fmt.Print("\n")
	fmt.Println("total", count)

	// 101 103 107 109 113 127 131 137 139 149 151 157 163 167 173 179 181 191 193 197 199
	// total 21
}
