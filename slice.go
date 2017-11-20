//切片Slice
//
//其本身并不是数组，它指向底层的数组
//作为变长数组的替代方案，可以关联底层数组的局部或全部
//为引用类型
//可以直接创建或从底层数组获取生成
//使用len()获取元素个数，cap()获取容量
//一般使用make()创建
//如果多个slice指向相同底层数组，其中一个的值改变会影响全部
//
//make([]T, len, cap)
//其中cap可以省略，则和len的值相同
//len表示存数的元素个数，cap表示容量
package main

import (
	"fmt"
)

func main() {

	//fmt.Println("slice开始")

	var s1 []int
	fmt.Println("刚刚申明的slice是:", s1)

	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("新申明的1到10的数组:", a)

	s2 := a[3]
	fmt.Println("索引位置3的位置值是:", s2)

	s3 := a[:5] //包含起始索引,不包含终止索引,这里表示的是0到4
	fmt.Println("索引位置0到4是:", s3)

	s4 := a[5:]
	fmt.Println("索引位置5到最后是:", s4)

	s5 := a[5:9]
	fmt.Println("索引位置5到8是", s5)

	//适用make声明slice
	s11 := make([]int, 3)
	fmt.Println("用make声明slice,没定义最大容量的情况", len(s11), cap(s11))

	s12 := make([]int, 3, 10)
	fmt.Println("用make声明slice,定义了最大容量的情况", len(s12), cap(s12))

	//Slice与底层数组的对应关系
	fmt.Println()
	fmt.Println("Slice与底层数组的对应关系")
	s13 := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := s13[2:5]
	fmt.Println(string(sa)) //cde
	sb := s13[3:5]
	fmt.Println(string(sb)) //de
	se := sa[2:3]
	fmt.Println(string(se)) //e
	sf := sa[2:5]
	fmt.Println(string(sf)) //efg
	//sg:=sa[2:100]
	//fmt.Println(string(sg))  //越界了,超过s13的最大容量位置

	fmt.Println()
	fmt.Println("使用append追加数组") //efg
	s21 := make([]int, 3, 6)
	fmt.Printf("%p\n", s21) //Printf

	s21 = append(s21, 1, 2, 3)
	fmt.Printf("%v %p\n", s21, s21)
	//第一次追加打印,没有超过容量地址没变

	s21 = append(s21, 1, 2, 3)
	fmt.Printf("%v %p\n", s21, s21)
	//第二次追加打印,超过了容量,地址发生了改变

	fmt.Println()
	fmt.Println("修改共享部分的数据,修改的一个slice,影响到另外一个slice的值")
	a101 := []int{1, 2, 3, 4, 5}
	s101 := a101[2:5] //3,4,5
	s102 := a101[1:3] //2,3
	fmt.Println(s101, s102)
	s101[0] = 888
	fmt.Println(s101, s102)

	fmt.Println("append超过slice容量将重新分配地址，修改原来的slice将不会修改到新分配的slice")
	a201 := []int{1, 2, 3, 4, 5}
	s201 := a201[2:5] //3,4,5
	s202 := a201[1:3] //2,3
	fmt.Println(s201, s202)
	s202 = append(s202, 4)
	s201[0] = 888
	fmt.Println(s201, s202)

	s202 = append(s202, 4, 555, 666, 777)
	s201[0] = 999
	fmt.Println(s201, s202)

	fmt.Println()
	fmt.Println("slice拷贝，以短的slice为准")
	s301 := []int{1, 2, 3, 4, 5, 6}
	s302 := []int{7, 8, 9}
	copy(s301, s302)
	fmt.Println("短的slice拷贝到长的slice中", s301) //789456

	s303 := []int{1, 2, 3, 4, 5, 6}
	s304 := []int{7, 8, 9}
	copy(s304, s303)
	fmt.Println("长的slice拷贝到短的slice中", s304) //123

	s305 := []int{1, 2, 3, 4, 5, 6}
	s306 := []int{7, 8, 9, 1, 1, 1, 1, 1, 1}
	copy(s306[2:4], s305[1:3]) //索引1和2的值是2,3 目标位置时s306的索引2,3
	fmt.Println("局部拷贝数据样例:", s306)

	g1 := []int{1, 2, 3, 4, 5}
	g2 := g1
	fmt.Println(g2) //[1,2,3,4,5]
	g3 := g1[0:5]
	fmt.Println(g3) //[1,2,3,4,5]
	g4 := g1[:]
	fmt.Println(g4) //[1,2,3,4,5]

}
