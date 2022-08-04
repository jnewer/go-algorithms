package main

import "fmt"

//有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
func main() {
	totalCount := 0
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				if i != k && i != j && j != k {
					totalCount++
					fmt.Println("第", totalCount, "方案", "i=", i, "j=", j, "k=", k)
				}
			}
		}
	}

	fmt.Println("共", totalCount, "种方案")
}
