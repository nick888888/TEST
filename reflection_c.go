//反射--修改1
package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("下面是用反射修改内容")
	x := 123
	fmt.Println(x)

	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)

	fmt.Println(x)
}
