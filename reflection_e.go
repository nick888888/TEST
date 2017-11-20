// 如何通过反射来对方法进行调用
// 那么让我们来写一个比较完整的通过反射对方法等动态调用

package main

import (
	"fmt"
	"reflect"
)

//定义一个用户结构体
type User struct {
	Id   int
	Name string
	Age  int
}

//为User绑定方法
func (u User) HelloDisplay(name string) {
	fmt.Println("Hello", name, " my name is ", u.Name)
}

func main() {
	u := User{1, "OK", 29}
	u.HelloDisplay("jack") //正常调用

	/**
	以下方式为反射调用，最优到代码写法就是新写一个方法且在开始是通过kind判断类型是否正确且需要判断有没有对应方法等
	*/
	v := reflect.ValueOf(u)                          //通过反射得到类型内容
	methodV := v.MethodByName("HelloDisplay")        //通过方法名称得道方法实体
	args := []reflect.Value{reflect.ValueOf("jack")} //设置反射传入的参数
	methodV.Call(args)
}
