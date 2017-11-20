//引用的,输出都是c
package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	select {}

}

//E:\go\data\src>go run pit_c1.go
//c
//c
//c
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [select (no cases)]:
//main.main()
//E:/go/data/src/pit_c1.go:14 +0x132
//exit status 2
