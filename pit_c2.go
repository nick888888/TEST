//引用的,输出都是c,要改变这种情况需要改成参数传递
package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	select {}
}

//E:\go\data\src>go run pit_c2.go
//g
//a
//b
//d
//f
//c
//e
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [select (no cases)]:
//main.main()
//E:/go/data/src/pit_c2.go:15 +0xad
//exit status 2
