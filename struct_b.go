//结构struct值拷贝,调用方法不会修改到原始值
package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {

	//字面值初始化
	a := person{
		Name: "joe",
		Age:  20,
	}
	fmt.Println("第一次打印的a:", a)
	A(a)
	fmt.Println("第二次打印的a:", a)
}

func A(per person) {
	per.Age = 13
	fmt.Println("A函数内的:", per)
}

//第一次打印的a: {joe 20}
//A函数内的: {joe 13}
//第二次打印的a: {joe 20}
