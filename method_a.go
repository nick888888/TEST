//方法method
//
//Go 中虽没有class，但依旧有method
//通过显示说明receiver来实现与某个类型的组合
//只能为同一个包中的类型定义方法
//Receiver 可以是类型的值或者指针
//不存在方法重载
//可以使用值或指针来调用方法，编译器会自动完成转换
//从某种意义上来说，方法是函数的语法糖，因为receiver其实就是
//方法所接收的第1个参数（Method Value vs. Method Expression）
//如果外部结构和嵌入结构存在同名方法，则优先调用外部结构的方法
//类型别名不会拥有底层类型所附带的方法
//方法可以调用结构中的非公开字段

package main

import (
	"fmt"
)

type A struct {
	B
	name string
}

type B struct {
	name string
}

func main() {
	a := A{}
	a.Print()

	b := B{}
	b.Print()
}

//方法绑定的对象类型不一样,就是不一样.即使名称是一样
func (a A) Print() { //接收者类型A
	fmt.Println("A方法打印")
}
func (b B) Print() { //接收者类型B
	fmt.Println("B方法打印")
}
