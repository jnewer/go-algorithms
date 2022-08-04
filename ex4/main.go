package main

import (
	"fmt"
)

//输入某年某月某日，判断这一天是这一年的第几天？
func main() {
	var y, m, d int = 0, 0, 0
	var days int = 0
	fmt.Printf("请输入年月日\n")
	fmt.Scanf("%d%d%d\n", &y, &m, &d)
	fmt.Printf("%d 年 %d 月 %d 日", y, m, d)
	switch m {
	case 12:
		days += d
		d = 30
		fallthrough
	case 11:
		days += d
		d = 31
		fallthrough
	case 10:
		days += d
		d = 30
		fallthrough
	case 9:
		days += d
		d = 31
		fallthrough
	case 8:
		days += d
		d = 31
		fallthrough
	case 7:
		days += d
		d = 30
		fallthrough
	case 6:
		days += d
		d = 31
		fallthrough
	case 5:
		days += d
		d = 30
		fallthrough
	case 4:
		days += d
		d = 31
		fallthrough
	case 3:
		days += d
		d = 28
		if (y%400 == 0) || (y%4 == 0 && y%100 != 0) {
			d += 1
		}
		fallthrough
	case 2:
		days += d
		d = 31
		fallthrough
	case 1:
		days += d
	}
	fmt.Printf("是今年的第 %d 天!\n", days)
}
