package main

import "fmt"

func main(){
	A()
	B()
	C()
	/*
	for i:=0; i<3; i++{
		defer func(){
			fmt.Println(i)
	}()
}
*/
	//fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")
}

func A(){
	fmt.Println("func A")
}

func B(){
	defer func(){
		if err := recover(); err != nil{
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
	
}

func C(){
	fmt.Println("func C")
}