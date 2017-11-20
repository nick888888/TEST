//并发concurrency
//
//很多人都是冲着 Go 大肆宣扬的高并发而忍不住跃跃欲试，但其实从
//源码的解析来看，goroutine 只是由官方实现的超级“线程池”而已。
//不过话说回来，每个实例 4-5KB 的栈内存占用和由于实现机制而大幅
//减少的创建和销毁开销，是制造 Go 号称的高并发的根本原因。另外，
//goroutine 的简单易用，也在语言层面上给予了开发者巨大的便利。
//
//并发不是并行：Concurrency Is Not Parallelism
//并发主要由切换时间片来实现“同时”运行，在并行则是直接利用
//多核实现多线程的运行，但 Go 可以设置使用核数，以发挥多核计算机
//的能力。
//
//Goroutine 奉行通过通信来共享内存，而不是共享内存来通信。

package main

import (
	"fmt"
	"time"
)

func main() {
	//启用一个goroutine
	go GoRun()
	//这里加一个休眠是因为主线程已启动就执行完毕消亡来，子线程还来不及执行
	fmt.Println("start sleep")
	time.Sleep(5 * time.Second)
	fmt.Println("end sleep")
	fmt.Println("test-git")

}

func GoRun() {
	fmt.Println("Go Go Go!!!")
}
