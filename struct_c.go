//通过指针的传递,加&取地址符号
package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("下面是指针传递修改结构person的值")
	a := person{
		Name: "joe",
		Age:  28,
	}
	//a.Name="ok"
	fmt.Println("第一次main中打印a:", a)
	A(&a)
	fmt.Println("第二次main中打印a:", a)
	B(&a)
	fmt.Println("第三次main中打印a:", a)
}

func A(per *person) {
	per.Age = 13
	fmt.Println("A FUNCTION", per)
}

func B(per *person) {
	per.Age = 15
	fmt.Println("B FUNCTION", per)
}

//E:\go\data\src>go run struct_c.go
//第一次main中打印a: {joe 28}
//A FUNCTION &{joe 13}
//第二次main中打印a: {joe 13}
//B FUNCTION &{joe 15}
//第三次main中打印a: {joe 15}
