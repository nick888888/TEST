//interface  接口
package main

import (
	"fmt"
)

//USB可以转换成Connect,但是Connect不能转换成USB
type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
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
	pc := PhoneConnecter{"PhoneConncter"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	//a.Name()
}
