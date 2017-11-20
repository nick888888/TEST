//数组Array

//定义数组的格式：var <varName> [n]<type>，n>=0
//数组长度也是类型的一部分，因此具有不同长度的数组为不同类型
//注意区分指向数组的指针和指针数组
//数组在Go中为值类型
//数组之间可以使用==或!=进行比较，但不可以使用<或>
//可以使用new来创建数组，此方法返回一个指向数组的指针
//Go支持多维数组
//Go语言版冒泡排序
package main

import (
	"fmt"
)

func main() {
		fmt.Println("git test111222")
	fmt.Println("git test")
	fmt.Println("下面是数组的常规声明")
	a1 := [2]int{1, 1}
	fmt.Println(a1)

	a2 := [2]int{1}
	fmt.Println(a2)

	a3 := [10]int{9: 1}
	fmt.Println(a3)

	a4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a4)

	a5 := [...]int{0: 1, 1: 2, 2: 3}
	fmt.Println(a5)

	var a6 [2]string
	fmt.Println(a6)

	a7 := [2]int{1, 2}
	fmt.Println(a7)

	a8 := [2]string{"11111111111", "hhh"}
	fmt.Println(a8)

	fmt.Println()
	fmt.Println("下面是指向数组的指针")
	//创建长度为10的数组
	c := [...]int{9: 1}
	fmt.Println(c)
	//取地址的符号,p是取数组的地址,指向数组的指针
	var p *[10]int = &c
	fmt.Println(p)
	//[0 0 0 0 0 0 0 0 0 1]
	//&[0 0 0 0 0 0 0 0 0 1]

	fmt.Println()
	fmt.Println("下面是指针数组")
	x, y := 1, 2
	d := [...]*int{&x, &y}
	fmt.Println(d) //输出的是变量x,y的地址,实际上保存的是指针
	//[0xc04200e130 0xc04200e138]

	fmt.Println()
	fmt.Println("使用new关键字返回的时指向数组的指针使用,不管是数组本身还是指向数组的指针,都可以用下面的方法赋值")
	b1 := [10]int{}
	b1[1] = 2
	fmt.Println(b1)
	p1 := new([10]int)
	p1[1] = 2
	fmt.Println(p1)

	fmt.Println()
	fmt.Println("下面是多维数组")
	e := [...][4]int{
		{1: 2},
		{3: 5}}
	fmt.Println(e)
	//[[0 2 0 0] [0 0 0 5]]

	fmt.Println()
	fmt.Println("下面是go的冒泡排序,从大到小排序")
	f := [...]int{5, 3, 4, 7, 8, 9}
	fmt.Println(f)

	num := len(f)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if f[i] < f[j] {
				var temp = f[i]
				f[i] = f[j]
				f[j] = temp
			}
		}
	}
	fmt.Println(f)

	fmt.Println()
	fmt.Println("下面是循环体内声明变量每次循环的实际地址发生变化样例")
	for i := 0; i < 3; i++ {
		v := 1
		fmt.Println("第", i, "次v的地址是:", &v)
	}

}
