//结构struct
//
//Go 中的struct与C中的struct非常相似，并且Go没有class
//使用 type <Name> struct{} 定义结构，名称遵循可见性规则
//支持指向自身的指针类型成员
//支持匿名结构，可用作成员或定义成员变量
//匿名结构也可以用于map的值
//可以使用字面值对结构进行初始化
//允许直接通过指针来读写结构成员
//相同类型的成员可进行直接拷贝赋值
//支持 == 与 !=比较运算符，但不支持 > 或 <
//支持匿名字段，本质上是定义了以某个类型名为名称的字段
//嵌入结构作为匿名字段看起来像继承，但不是继承
//可以使用匿名字段指针

package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {

	//最简单的初始化和定义
	a1 := person{}
	a1.Name = "jan"
	a1.Age = 19
	fmt.Println(a1) //{jan 19}

	//字面值初始化,这个比较方便,经常使用
	a2 := person{
		Name: "air",
		Age:  88,
	}
	fmt.Println(a2)

}
