//指针传递,可以修改前面的值
package main

import (
	"fmt"
)

func main() {
	a := 1
	A(&a) //取出a的地址
	fmt.Println(a)
}

func A(a *int) { //指针类型传递
	*a = 2
	fmt.Println(*a)
}

//2
//2
