//Go 没有异常机制，但有 panic/recover 模式来处理错误
//Panic 可以在任何地方引发，但recover只有在defer调用的函数中有效

package main

import (
	"fmt"
)

func main() {

	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

//defer 放在panic之前才能恢复,panic是直接结束了.必须提前告知有defer才可能恢复程序代码.

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}

//E:\go\data\src>go run FUNCTION_j_recover_panic.go
