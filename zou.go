package main

import "fmt"

func main() {
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	// 接着我们在源 slice 之上创建一个新的 slice

	slice := source[2:3:4]
	fmt.Println(slice)

}
