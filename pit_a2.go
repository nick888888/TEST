//append后地址重新分配了
package main

import (
	"fmt"
)

func Pingpong(s []int) []int { //设计一个返回值
	s = append(s, 3)
	return s
}

func main() {
	s := make([]int, 0) //初始容量append不够用
	fmt.Println(s)
	s = Pingpong(s)
	fmt.Println(s)
}

//E:\go\data\src>go run pit_a2.go
//[]
//[3]
