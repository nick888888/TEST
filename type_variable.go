//类型与变量
//Type and variable

//Go基本类型

//- 布尔型：bool
//  长度：1字节
//  取值范围：true, false
//  注意事项：不可以用数字代表true或false

//- 整型：int/uint
//  根据运行平台可能为32或64位

//- 8位整型：int8/uint8
//  长度：1字节
//  取值范围：-128~127/0~255
//  字节型：byte（uint8别名）

//- 16位整型：int16/uint16
//  长度：2字节
//  取值范围：-32768~32767/0~65535
//- 32位整型：int32（rune）/uint32
//  长度：4字节
//  取值范围：-2^32/2~2^32/2-1/0~2^32-1
//- 64位整型：int64/uint64
//  长度：8字节
//  取值范围：-2^64/2~2^64/2-1/0~2^64-1
//- 浮点型：float32/float64
//  长度：4/8字节
//  小数位：精确到7/15小数位

//- 复数：complex64/complex128
//  长度：8/16字节
//  足够保存指针的 32 位或 64 位整数型：uintptr

//- 其它值类型：
//  array、struct、string
//- 引用类型：
//  slice、map、chan

//- 接口类型：inteface
//- 函数类型：func

package main

import (
	"fmt"
	"math"
	"strconv"
)

type (
	byte     int8
	rune     int32
	ByteSize int64
	//别名
	文本 int16
)

func main() {

	//0值
	var a1 int
	fmt.Println(a1)

	//0
	var a2 int32
	fmt.Println(a2)

	//0
	var a3 float32
	fmt.Println(a3)

	var a4 bool
	fmt.Println(a4)

	var a5 string
	fmt.Println(a5)

	var a6 [2]int
	fmt.Println(a6)

	var a7 []int
	fmt.Println(a7)

	var a8 [3]bool
	fmt.Println(a8)

	var a9 [2]byte
	fmt.Println(a9)

	//
	var a10 文本
	fmt.Println(a10 + 55)

	//检查
	fmt.Println(math.MinInt8)

	b := false
	fmt.Println(b)

	//多个变量的声明与赋值
	//
	//全局变量的声明可使用 var() 的方式进行简写
	//全局变量的声明不可以省略 var，但可使用并行方式
	//所有变量都可以使用类型推断
	//局部变量不可以使用 var() 的方式简写，只能使用并行方式

	var z, x, _, c = 1, 2, 3, 4 //空白符号，忽略，应用的多个函数返回值中
	fmt.Println(z)
	fmt.Println(x)
	fmt.Println(c)

	//变量的类型转换
	//
	//Go中不存在隐式转换，所有类型转换必须显式声明
	//转换只能发生在两种相互兼容的类型之间
	//类型转换的格式：
	//<ValueA> [:]= <TypeOfValueA>(<ValueB>)
	var a111 float32 = 100.1
	fmt.Println(a111)
	b111 := int(a111)
	fmt.Println(b111)

	//
	var d int = 65
	fmt.Println(d) //65

	f := string(d)
	fmt.Println(f) //A

	var d1 int = 65
	f1 := strconv.Itoa(d1)
	fmt.Println(f1) //65

	f2, _ := strconv.Atoi(f1)
	fmt.Println(f2) //65

}
