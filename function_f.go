//函数也可以作为参数传递,go一切皆类型
package main

import (
	"fmt"
)

func main() {
	a := A
	a()
}

func A() {
	fmt.Print("FUNC A")
}
