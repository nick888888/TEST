//值拷贝,不会改变前面的值
package main

import (
	"fmt"
)

func main() {
	a := 1
	A(a)
	fmt.Println(a)
}

func A(a int) {
	a = 2
	fmt.Println(a)
}

//2
//1
