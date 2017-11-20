//值类型传递和指针类型传递

package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Name = "a.Name未修改前的值"
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Name = "b.Name未修改前的值"
	b.Print()
	fmt.Println(b.Name)
}

func (a *A) Print() {
	a.Name = "a.Name方法中的值"
	fmt.Println("运行了A的Print")
}

func (b B) Print() {
	b.Name = "b.Name方法中的值"
	fmt.Println("运行了B的Print")
}
