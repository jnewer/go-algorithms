package main

import "fmt"

// 有 5 个人坐在一起，问第五个人多少岁？他说比第 4 个人大 2 岁。
// 问第 4 个人岁数，他说比第 3 个人大 2 岁。问第三个人，又说比第 2 人大两岁。
// 问第 2 个人，说比第一个人大两岁。最后问第一个人，他说是 10 岁。请问第五个人多大？
func calcAges(n int) (age int) {
	if n == 1 {
		age = 10
	} else {
		age = calcAges(n-1) + 2
	}
	return
}
func main() {
	fmt.Printf("第五个人年龄为：%d\n", calcAges(5)) //第五个人年龄为：18
}
