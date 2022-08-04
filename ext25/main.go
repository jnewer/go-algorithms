package main

import "fmt"

// 求 1+2!+3!+…+20! 的和。
func main() {
	s, t := 0, 1
	for n := 1; n <= 20; n++ {
		t *= n
		s += t
	}
	fmt.Printf("1+2!+3!...+20!=%d\n", s) //1+2!+3!...+20!=2561327494111820313
}
