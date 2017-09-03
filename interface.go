package main

import "fmt"

type USB interface{
	Name() string
	Connecter//借口的嵌套
}

type Connecter interface{
	Connect()
}

type PhoneConnecter struct{
	name string
}

type TVConnecter struct{
	name string
}

func (tv TVConnecter)Connect(){
	fmt.Println("TVConnnect:" , tv.name)
}

func (pc PhoneConnecter)Name() string{
	return pc.name
}

func (pc PhoneConnecter) Connect(){
	fmt.Println("Connect: ", pc.name)
}
/*
func Disconnect(usb USB){
	fmt.Println("Disconnect...")
	if pc, ok := usb.(PhoneConnecter);ok{
		fmt.Println("Disconnected: ", pc.name)
		return
	}
	fmt.Println("unknown device.")
}
*/

func Disconnect(usb interface{}){
	// 通过类型断言的ok pattern可以判断接口中的数据类型

	/*
	fmt.Println("Disconnect...")
	if pc, ok := usb.(PhoneConnecter);ok{
		fmt.Println("Disconnected: ", pc.name)
		return
	}
	*/

	// 使用type switch则可针对空接口进行比较全面的类型判断
	switch v:= usb.(type){  //传入的是空接口
		case PhoneConnecter:
			fmt.Println("Disconnected:", v.name)
		default:
			fmt.Println("Unknown device...")
	}
	fmt.Println("unknown device.")
}


func main(){
	/*
	tv := TVConnecter{"TVConnecter"}
	var a USB
	a = USB(tv)
	a.Connect()
	*/
	//var a USB
	
	pc := PhoneConnecter{"Phone123"}
	var a Connecter
	a = Connecter(pc)
	//a := PhoneConnecter{"Phone"}
	a.Connect()
	
	pc.name = "PC"
	a.Connect()

	var b interface{}
	fmt.Println(b == nil)

	var p *int = nil
	b = p
	fmt.Println(b == nil)
	
	//Disconnect(a)
}