package main

import "fmt"


type A struct{
	Name string
}

type B struct{
	Name string
}

func main(){
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)
}

func (a *A) Print(){
	a.Name = "AA"
	fmt.Println("A")
}

// 不支持函数方法的重载,所以这是不能同时存在的
/*
func (a A)Print(b int){

}
*/

func (b B) Print(){
	b.Name = "BB"
	fmt.Println("B")
}