//从某种意义上来说，方法是函数的语法糖，因为receiver其实就是
//方法所接收的第1个参数（Method Value vs. Method Expression）

package main

import (
	"fmt"
)

type TZ int

func main() {
	var a TZ
	a.Print()       //Method Value 方法值
	(*TZ).Print(&a) //Method Expression方法表达式
}

func (a *TZ) Print() {
	fmt.Println("TZ－t3333")
}
