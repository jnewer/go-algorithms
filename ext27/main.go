package main

import "fmt"

func putchar(n int) {
	var c byte
	if n >= 1 {
		fmt.Scanf("%c", &c)
		putchar(n - 1)
		fmt.Printf("%c", c)
	}
}
func main() {
	fmt.Printf("请输入%d个字符：", 5)
	putchar(5)
}
