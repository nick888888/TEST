//匿名结构1

package main

import (
	"fmt"
)

func main() {
	a1 := struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  30,
	}
	fmt.Println(a1)

}
