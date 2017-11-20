//判断语句if
//
//条件表达式没有括号
//支持一个初始化表达式（可以是并行方式）
//左大括号必须和条件语句或else在同一行
//支持单行模式
//初始化语句中的变量为block级别，同时隐藏外部同名变量
//1.0.3版本中的编译器BUG

package main

import (
	"fmt"
)

func main() {

	//执行if语句的时候，外部同名变量会隐藏起来不执行，适用if语句块中的变量执行，if语句块中的变量的作用域在if
	//语句块中。
	a1 := 10
	if a1 := 1; a1 > 0 {
		fmt.Println(a1)
	}
	fmt.Println(a1)

	//循环语句for
	//Go只有for一个循环语句关键字，但支持3种形式
	//初始化和步进表达式可以是多个值
	//条件语句每次循环都会被重新检查，因此不建议在条件语句中
	//使用函数，尽量提前计算好条件并以变量或常量代替
	//左大括号必须和条件语句在同一行
	fmt.Println()
	fmt.Println("xiamianshi for1")
	a2 := 1
	for {
		a2++
		if a2 > 3 {
			break
		}
		fmt.Println("if hou", a2)
	}
	fmt.Println("for hou", a2)
	fmt.Println("OVER for1")

	fmt.Println()
	fmt.Println("xiamianshi for2")
	a3 := 1
	for a3 <= 3 {
		a3++
		fmt.Println(a3)
	}
	fmt.Println("OVER for2")

	fmt.Println()
	fmt.Println("xiamianshi for3")
	a4 := 1
	for i := 0; i < 3; i++ {
		a4++
		fmt.Println("for nei", a4)
	}
	fmt.Println("OVER for3")

	//选择语句switch
	//可以使用任何类型或表达式作为条件语句
	//不需要写break，一旦条件符合自动终止
	//如希望继续执行下一个case，需使用fallthrough语句
	//支持一个初始化表达式（可以是并行方式），右侧需跟分号
	//左大括号必须和条件语句在同一行

	//形式1，这里是放常量 输出结果是  a5=1
	fmt.Println()
	a5 := 1
	switch a5 {
	case 0:
		fmt.Println("a5=0")
	case 1:
		fmt.Println("a5=1")
	default:
		fmt.Println("None")
	}

	//形式2,这里输出a6>0,并没有输出a6>1
	fmt.Println()
	a6 := 1
	switch {
	case a6 > 0:
		fmt.Println("a6>0")
	case a6 > 1:
		fmt.Println("a6>1")
	default:
		fmt.Println("None")
	}

	//形式3,这里输出a7>0,a7>1
	fmt.Println()
	a7 := 1
	switch {
	case a7 > 0:
		fmt.Println("a7>0")
		fallthrough
	case a7 > 1:
		fmt.Println("a7>1")
	default:
		fmt.Println("None")
	}

	//形式4,这里输出a7>0,a7>1  超出了switch,定义的a8就不能用了
	fmt.Println()
	switch a8 := 1; {
	case a8 > 0:
		fmt.Println("a8>0")
		fallthrough
	case a8 > 1:
		fmt.Println("a8>1")
	default:
		fmt.Println("None")
	}
	//fmt.Println(a8)   //undefined: a8

	//跳转语句goto, break, continue
	//
	//三个语法都可以配合标签使用
	//标签名区分大小写，若不使用会造成编译错误
	//Break与continue配合标签可用于多层循环的跳出
	//Goto是调整执行位置，与其它2个语句配合标签的结果并不相同
	fmt.Println("下面是跳转语句goto, break, continue")
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1 //break 跳出与LABEL1同级别的循环
			}
		}
	}
	fmt.Println("break-OK")

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL2 //直接调整程序执行的位置
			}
		}
	}
LABEL2:
	fmt.Println("goto-OK")

LABEL3:
	for i := 0; i <= 10; i++ {
		for {
			fmt.Println("进入的i:", i)
			continue LABEL3 //中止当前循环,执行下一次循环.
			fmt.Println("永远执行不到的代码")
		}
	}
	fmt.Println("continue-OK")

}
