package main

import "fmt"

type B bool

func (b B) C(t, f interface{}) interface{} {
	if bool(b) == true {
		return t
	}
	return f
}

//利用条件运算符的嵌套来完成此题：学习成绩 >= 90 分的同学用 A 表示，60-89 分之间的用 B 表示，60 分以下的用 C 表示。
func main() {
	var score int = 0
	fmt.Print("输入分数：")
	fmt.Scanf("%d", &score)
	grade := B(score >= 90).C("A", B(score >= 60).C("B", "C"))
	fmt.Println("grade = ", grade)
}
