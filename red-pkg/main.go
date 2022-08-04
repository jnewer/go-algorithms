package main

import (
	"fmt"
	"math/rand"
	"time"
)

// remainCount: 剩余红包数
// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount, remainMoney int) int {
	if remainCount == 1 {
		return remainMoney
	}

	rand.Seed(time.Now().UnixNano())

	var min = 1
	max := remainMoney / remainCount * 2
	money := rand.Intn(max) + min
	return money

}

// count: 红包数量
// money: 红包金额（单位：分)
func redPackage(count, money int) {
	for i := 0; i < count; i++ {
		m := randomMoney(count-i, money)
		fmt.Printf("%d ", m)
		money -= m
	}
}

//抢红包算法即类似微信拼手气红包，发一定金额的红包，指定人数抢红包
func main() {
	for i := 0; i < 10; i++ {
		redPackage(10, 500)
		fmt.Println("")
	}
}
