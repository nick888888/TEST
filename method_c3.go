//方法可以调用结构中的非公开字段

package main

import (
	"fmt"
)

type A struct {
	name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println("main中", a.name)
}

func (a *A) Print() {
	a.name = "1111"
	fmt.Println(a.name)
}

//E:\go\data\src>go run method_C3.go
//1111
