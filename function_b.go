//函数虽然接收的时一个slice,但是还是没改变以前的值,其实是值拷贝.
package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	A(a, b)
	fmt.Println(a, b)

}

func A(s ...int) {
	s[0] = 5
	s[1] = 6
	fmt.Println(s)
}

//[5 6]
//1 2
