//嵌入结构

package main

import (
	"fmt"
)

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {

	fmt.Println("嵌入结构")
	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := teacher{Name: "joe", Age: 20, human: human{Sex: 1}}
	a.Name = "joe2"
	a.Age = 55
	a.Sex = 100
	fmt.Println(a, b)

}

//E:\go\data\src>go run struct_g.go
//嵌入结构
//{{100} joe2 55} {{1} joe 20}
