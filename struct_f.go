//匿名结构2  初始化赋值

package main

import (
	"fmt"
)

type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone, City string
	}
}

type person3 struct {
	string
	int
}

func main() {
	//匿名结构 初始化赋值
	a2 := person{Name: "jeo", Age: 19}
	a2.Contact.Phone = "12343435"
	a2.Contact.City = "beijing"
	fmt.Println(a2)

	//匿名字段 一定要按照顺序来赋值
	a3 := person3{"jjj", 33}
	fmt.Println("a3:", a3)

	var a4 person3
	a4 = a3
	fmt.Println("a4:", a4)

}
