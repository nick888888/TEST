//匿名函数
package main

import (
	"fmt"
)

func main() {
	a := func() { //匿名函数
		fmt.Print("FUNC A")
	}
	a()
}
