//课堂作业

//创建一个 goroutine，与主线程按顺序相互发送信息若干次并打印

package main

import (
	"fmt"
)

var c chan string

func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Pingpong:hi,#%d", i)
		i++
	}
}

func main() {
	c = make(chan string) //初始化chan
	go Pingpong()
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("From main:Hello,#%d", i)
		fmt.Println(<-c)
	}
}
