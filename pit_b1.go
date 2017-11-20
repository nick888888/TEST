//时间格式一定要用指定的常量
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)                    //2017-11-16 13:59:40.91 +0800 CST m=+0.006000001
	fmt.Println(t.Format(time.ANSIC)) //Thu Nov 16 14:01:28 2017
	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(t.Format("Thu Nov 16 14:01:28 2017"))
	//E:\go\data\src>go run pit_b1.go
	//2017-11-16 14:03:29.373 +0800 CST m=+0.003000001
	//Thu Nov 16 14:03:29 2017
	//Thu Nov 16 14:03:29 2017
	//Thu Nov 116 113:11:168 16117
}
