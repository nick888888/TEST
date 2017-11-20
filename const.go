//常量的初始化规则与枚举
//
//在定义常量组时，如果不提供初始值，则表示将使用上行的表达式
//使用相同的表达式不代表具有相同的值
//iota是常量的计数器，从0开始，组中每定义1个常量自动递增1
//通过初始化规则与iota可以达到枚举的效果
//每遇到一个const关键字，iota就会重置为0

package main

import (
	"fmt"
)

//大写开头会被包的外部使用到。下划线和小写的c
const (
	a = "123"
	b = len(a)
	c
	d, e = 111, 222
	f, g
)
const (
	a1 = 'A'
	b1
	c1 = iota //与定义的顺序有关系，与出现的次数无关。
	d1
)
const (
	e1 = iota //0
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(a1) //65
	fmt.Println(b1) //65
	fmt.Println(c1) //2
	fmt.Println(d1) //3

	fmt.Println(e1) //0
}
