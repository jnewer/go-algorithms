package main

import (
	"bufio"
	"fmt"
	"os"
)

//输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
func main() {
	var i, j, k, l = 0, 0, 0, 0
	fmt.Print("请输入一串字符：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	for _, ch := range input {
		switch {
		case ch >= 'A' && ch <= 'Z':
			i++
		case ch > 'a' && ch <= 'z':
			i++
		case ch == ' ' || ch == '\t':
			j++
		case ch >= '0' && ch < '9':
			k++
		default:
			l++
		}
	}

	fmt.Printf("char count = %d, space count = %d, digit count = %d, others count = %d\n", i, j, k, l)
}
