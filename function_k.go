package main

import (
	"fmt"
)

func main() {

	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		//这的i是一个参数,值拷贝,应该是0123,由于是defer定义的所以是逆序输出3210
		defer fmt.Println("defer i=", i)

		//也是用到闭包的思想,但是这里都是4,所以输出都是4
		defer func() { fmt.Println("defer_closure i=", i) }()

		//匿名函数,没有参数,直接使用了i,并没有定义i,是从当前匿名函数的外层读取
		//到i这个地址的引用地址,执行完上面的for循环的时候这个地址的值是4,所以输出的都是4
		fs[i] = func() { fmt.Println("closure i=", i) }

	}
	for _, f := range fs {
		f()
	}
}

//E:\go\data\src>go run FUNCTION_k.go
//closure i= 4
//closure i= 4
//closure i= 4
//closure i= 4
//defer_closure i= 4
//defer i= 3
//defer_closure i= 4
//defer i= 2
//defer_closure i= 4
//defer i= 1
//defer_closure i= 4
//defer i= 0
