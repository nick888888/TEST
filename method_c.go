//类型别名和方法的组合,强制类型转换.
//package包内,方法可以访问结构的私有字段

package main

import (
	"fmt"
)

type TZ int

func main() {
	var a TZ
	a.Print()
	var b TZ

	b = 66
	fmt.Print(b)
}

func (a *TZ) Print() {
	fmt.Println("TZ－t")
}
