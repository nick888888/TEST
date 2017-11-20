//嵌入结构
package main

import (
	"fmt"
)

type A struct {
	B
	Name string
}

type B struct {
	Name string
}

func main() {

	fmt.Println("嵌入结构")
	a := A{Name: "A", B: B{Name: "B"}}
	fmt.Println(a.Name, a.B.Name)
	//最高级别先找,每找到往下一级别找,要是下一级别有多个,就会提醒同名无法选择了
}
