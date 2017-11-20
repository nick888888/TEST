//使用 for range 来进行消息的发送

package main

import (
	"fmt"
)

func main() {
	//声明创建一个通道，存储类型为bool型，这里设置的channel就是双向通道，既可以存也可以取
	c := make(chan bool)
	//启用一个goroutine,使用的是匿名方法方式
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true //向 channel 中存入一个值
		close(c)  //切记如果使用for range来进行取值的时候需要在某个地方进行关闭，否则会发生死锁
	}()
	//从通道中循环取出刚才赋的值
	for v := range c {
		fmt.Println(v)
	}
}
