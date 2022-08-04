package main

import "fmt"

//求 s = a + aa + aaa + aaaa + aa…a 的值，其中 a 是一个数字。
//例如 2+22+222+2222+22222(此时共有 5 个数相加)，几个数相加由键盘控制。
func main() {
	var a, n, count int
	var sum, tn int
	fmt.Printf("please input a and n : ")
	fmt.Scanf("%d%d\n", &a, &n)
	for count < n {
		tn = tn + a
		sum = sum + tn
		a = a * 10
		count++
	}

	fmt.Printf("Sum = %d\n", sum)
}
