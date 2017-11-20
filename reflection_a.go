//反射获取字段,字段的值,判断传递进去的类型
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

func (u User) Hello() {
	fmt.Println("Hello world")
}

func main() {
	u := User{1, "ok", 12}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//检验对象类型,判断,不等于错误提示.
	if k := t.kind(); k != reflect.Struct {
		fmt.Println("xx")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

}

//E:\go\data\src>go run reflection.go

//Type: User
//Fields:

//Id: int = 1
//Name: string = ok
//Age: int = 12

//Hello: func(main.User)
