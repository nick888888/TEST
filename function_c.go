//这个是地址拷贝,会改变前面的值,
package main

import (
	"fmt"
)

func main() {
	s1 := []int{1, 2, 3, 4}
	A(s1)
	fmt.Println(s1)
}

func A(s []int) {
	s[0] = 8
	s[1] = 9
	s[2] = 10
	s[3] = 11
	fmt.Println(s)
}

//[8 9 10 11]
//[8 9 10 11]
