//interface  嵌入接口,类型断言

package main

import (
	"fmt"
)

//定义接口,嵌入接口
type USB interface {
	Name() string
	Connect
}

type Connect interface {
	Connect()
}

type empty interface {
}

//实现接口
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter name"}
	a.Connect()
	Disconnect(a)
	fmt.Println("下面是用switch type判断")
	Disconnect2(a)
}

//简单的类型断言ok pattern
func Disconnect(usb USB) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected.", pc.name)
		return
	}
	fmt.Println("Unknown decive.")

}

//switch
//方法多的接口可以向方法少的接口转化

func Disconnect2(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}

}
