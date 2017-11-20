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
	a := &person{ //初始化就取地址
		Name: "joe",
		Age:  15,
	}
	a.Name = "ok" //这里直接使用a,不用使用&a取地址
	fmt.Println("第一次main中打印a:", a)
	A(a)
	fmt.Println("第二次main中打印a:", a)
	B(a)
	fmt.Println("第三次main中打印a:", a)
}

func A(per *person) { //只需要在接收参数的时候加*
	per.Age = 20
	fmt.Println("A FUNCTION", per)
}

func B(per *person) {
	per.Age = 25
	fmt.Println("B FUNCTION", per)
}

//E:\go\data\src>go run struct_d.go
//下面是指针传递修改结构person的值
//第一次main中打印a: &{joe 15}
//A FUNCTION &{joe 20}
//第二次main中打印a: &{joe 20}
//B FUNCTION &{joe 25}
//第三次main中打印a: &{joe 25}
