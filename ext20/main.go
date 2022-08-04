package main

import "fmt"

//一球从100 米高度自由落下，每次落地后反跳回原高度的一半；再落下，求它在第 10 次落地时，共经过多少米？第 10 次反弹多高？
func main() {
	s := 100.0
	h := s / 2
	for i := 2; i < 10; i++ {
		s += 2 * h
		h /= 2
	}

	fmt.Printf("总距离：%f\n", s)
	fmt.Printf("最后一次高度：%f\n", h)
}
