//闭包
package main

import (
	"fmt"
)

func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

//0xc0420080c8
//0xc0420080c8
//11
//0xc0420080c8
//12
