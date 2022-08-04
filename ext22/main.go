package main

import "fmt"

// 两个乒乓球队进行比赛，各出三人。甲队为 a、b、c 三人，乙队为x、y、z 三人。
// 已抽签决定比赛名单。有人向队员打听比赛的名单。
// a 说他不和 x 比，c 说他不和 x、z 比，请编程序找出三队赛手的名单。
func main() {
	var i, j, k int16 /*i 是 a 的对手，j 是 b 的对手，k 是 c 的对手*/
	for i = 'x'; i <= 'z'; i++ {
		for j = 'x'; j <= 'z'; j++ {
			if i != j {
				for k = 'x'; k <= 'z'; k++ {
					if i != k && j != k {
						if i != 'x' && k != 'x' && k != 'z' {
							fmt.Printf("比赛对手： a--%c b--%c c--%c\n", i, j, k) //比赛对手： a--z b--x c--y
						}
					}
				}
			}
		}
	}
}
