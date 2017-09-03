package main

import "fmt"

type TZ int //可以为type命名的类型别名添加方法

func main(){
	var a TZ
	a.Print()
	(*TZ).Print(&a)
}

func (a *TZ)Print(){
	fmt.Println("TZ")
}