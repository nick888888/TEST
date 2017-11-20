//反射匿名和嵌入字段的问题
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func main() {
	fmt.Println("下面是匿名Anonymous")
	m := Manager{User: User{1, "myname", 33}, title: "122"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))
	//reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x4b7500), Tag:"", Offset:0x0, Index:[
	//]int{0}, Anonymous:true}

	fmt.Println()

	fmt.Printf("%#v\n", t.Field(1))
	//reflect.StructField{Name:"title", PkgPath:"main", Type:(*reflect.rtype)(0x4a8d60), Tag:"", Offset:0x20, I
	//	ndex:[]int{1}, Anonymous:false}

	fmt.Println()
	fmt.Println("下面是如何取到User中的Name,Id,Age,用索引查找")
	//如何取到User中的Name,Id,Age,用索引查找 传递一个int型的slice
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	//reflect.StructField{Name:"Id", PkgPath:"", Type:(*reflect.rtype)(0x4a8900), Tag:"", Offset:0x0, Index:[]i
	//	nt{0}, Anonymous:false}

}
