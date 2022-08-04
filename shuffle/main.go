package main

import (
	"fmt"
	"math/rand"
	"time"
)

func shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	var i, j int
	var temp int
	for i = len(arr) - 1; i > 0; i-- {
		j = rand.Intn(i + 1)
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
}

// 洗牌算法，即将原来的顺序打乱，组成新的随机排序的顺序。
func main() {
	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10; i++ {
		shuffle(intArr)
		fmt.Println(intArr)
	}
}
