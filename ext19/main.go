package main

import "fmt"

//一个数如果恰好等于它的因子之和，这个数就称为 “完数”。例如 6=1＋2＋3，编程找出 1000 以内的所有完数。
func main() {
	for n := 2; n < 1000; n++ {
		m := n
		for i := 1; i < n; i++ {
			if n%i == 0 {
				m -= i
			}
		}
		if m == 0 {
			fmt.Println("Num is: ", n)
		}
	}
}
