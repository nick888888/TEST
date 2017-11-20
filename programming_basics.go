//当前程序的包名
package main

//导入其他的包
import (
	"fmt"
	//别名,省略调用
)

//import  std  "fmt"
//import  .  "fmt"

//常量的定义
const PI = 3.14

const (
	PIO = 654
	PIR = 666
)

//全局变量的声明与赋值
var name = "gopher"
var (
	name1 = "1"
	name2 = 2
	name3 = 3
)

//一般类型的声明
type newType int
type (
	type1 int
	type2 string
	type3 bool
)

//结构的声明
type structName struct{}

//接口的声明
type golang interface{}

//程序入口
func main() {
	fmt.Println("hello world!你好")
}
