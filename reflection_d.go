//反射--使用反射修改接口中的值(那么让我们来写一个比较完整的通过反射修改
// 结构体内部字段内容)
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

func main() {
	fmt.Println("下面是用反射修改内容")
	u := User{1, "egg", 33}
	fmt.Println("设置前的:", u)
	Set(&u)
	fmt.Println("设置后的:", u)
}

func Set(o interface{}) { //空接口
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //判断传递进来的是不是
		fmt.Println("XXX")
		return
	} else {
		v = v.Elem() //两个条件都满足,赋值
	}

	f1 := v.FieldByName("Name") //用字段名称来返回值
	if !f1.IsValid() {          //如果没有找到,返回一个BAD
		fmt.Println("没找到名字Name")
		return
	}
	if f1.Kind() == reflect.String {
		f1.SetString("BYEBYE")
	}

	f2 := v.FieldByName("Name1") //用字段名称来返回值
	if !f2.IsValid() {           //如果没有找到,返回一个BAD
		fmt.Println("没找到名字Name1")
		return
	}
	if f2.Kind() == reflect.String {
		f2.SetString("1111")
	}

}
