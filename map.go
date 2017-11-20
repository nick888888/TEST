//map

//类似其它语言中的哈希表或者字典，以key-value形式存储数据
//Key必须是支持==或!=比较运算的类型，不可以是函数、map或slice
//Map查找比线性搜索快很多，但比使用索引访问数据的类型慢100倍
//Map使用make()创建，支持 := 这种简写方式

//make([keyType]valueType, cap)，cap表示容量，可省略
//超出容量时会自动扩容，但尽量提供一个合理的初始值
//使用len()获取元素个数

//键值对不存在时自动添加，使用delete()删除某键值对
//使用 for range 对map和slice进行迭代操作

package main

import (
	"fmt"
	"sort"
)

func main() {

	//map的初始化
	fmt.Println("map的初始化")
	var m1 map[int]string
	m1 = map[int]string{}
	fmt.Println("m1是", m1)

	var m2 map[int]string
	m2 = make(map[int]string)
	fmt.Println("m2是", m2)

	var m3 map[int]string = make(map[int]string)
	fmt.Println("m3是", m3)

	m4 := make(map[int]string)
	fmt.Println("m4是", m4)

	//map的存取
	fmt.Println()
	fmt.Println("map的存取")
	m5 := make(map[int]string)
	m5[1] = "ok5"
	fmt.Println("map的存入的值:", m5[1])
	a5 := m5[1]
	fmt.Println("map的取出的值:", a5)

	//map的删除
	fmt.Println()
	fmt.Println("map的删除")
	m6 := make(map[int]string)
	m6[1] = "ok6"
	fmt.Println("存入后的值:", m6[1])
	delete(m6, 1)
	fmt.Println("删除后的值:", m6[1])

	//复杂的map,每一级map都要有单独的初始化
	fmt.Println()
	fmt.Println("两层map")
	var m7 map[int]map[int]string
	m7 = make(map[int]map[int]string)
	m7[1] = make(map[int]string)
	//使用1只是对1进行了初始化.
	m7[1][1] = "ok7"
	a7 := m7[1][1]
	fmt.Println(a7)

	//map多返回值
	fmt.Println()
	fmt.Println("map多返回值情况,告知键值对是否存在,true表示存在")
	var m8 map[int]map[int]string
	m8 = make(map[int]map[int]string)
	a8, ok := m8[2][1] //map多返回值
	//检查,如果不存在需要make
	if !ok {
		m8[2] = make(map[int]string)
	}
	m8[2][1] = "ok888"
	a8 = m8[2][1]
	fmt.Println("取了第一次ok", a8, ok) //ok888 false

	a8, ok = m8[2][1]
	fmt.Println("取了第二次ok", a8, ok) //ok888 true

	// //迭代操作,使用 for range 对map和slice进行迭代操作
	// /*   for i,v:=range slice{   //
	// 	slice[i]=8
	// }  */

	//下面i是索引,v是值,修改v的值不会影响到slice本身,如果需要修改slice本身,用k进行操作
	//for i,v:=range slice{
	//
	//}
	//
	//for k,v:=range map{
	//
	//}

	//修改V,不会改到sm本身.得到的v是拷贝
	sm := make([]map[int]string, 5) //5是slice的长度
	for _, v := range sm {
		v = make(map[int]string)
		//v = make(map[int]string,1)  //对slice中的map进行初始化,后面这个参数的作用?
		v[2] = "ok"
		fmt.Println(v)
	}
	fmt.Println(sm)

	//	map[2:ok]
	//	map[2:ok]
	//	map[2:ok]
	//	map[2:ok]
	//	map[2:ok]
	//  [map[] map[] map[] map[] map[]]

	//需要对被迭代对象进行操作的话,要用i,k等等操作
	//对map进进行间接排序没
	//

	map1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h"}
	slice1 := make([]int, len(map1))
	i := 0
	for k, _ := range map1 {
		slice1[i] = k
		i++
	}
	fmt.Println("排序前的slice:", slice1)
	sort.Ints(slice1)
	fmt.Println("排序后的slice:", slice1)
	//排序前的slice: [1 3 4 8 2 5 6 7]
	//排序后的slice: [1 2 3 4 5 6 7 8]

	fmt.Println()
	fmt.Println("下面是k,v")
	//根据在 for range 部分讲解的知识，尝试将类型为map[int]string
	//的键和值进行交换，变成类型map[string]int
	//程序正确运行后应输出如下结果：
	map11 := map[int]string{0: "j", 8: "i", 5: "g", 4: "k", 3: "p"}
	fmt.Println(map11)
	map22 := make(map[string]int)
	for k, v := range map11 {
		map22[v] = k
	}
	fmt.Println(map22)
	//显示的结果
	//map[0:j 8:i 5:g 4:k 3:p]
	//map[i:8 g:5 k:4 p:3 j:0]对调
}
