package main

import "fmt"

type person struct{
	Name string
	Age int
}

func main(){
	a := &person{
		Name: "john",
		Age: 19,
	}
	a.Name = "ok"
	//a.Name = "zyh"
	//a.Age = 19
	fmt.Println(a)
	A(a)
	B(a)
	fmt.Println(a)
}

func A(per *person){
	per.Age = 13
	fmt.Println("A", per)
}

func B(per *person){
	per.Age = 15
	fmt.Println("B", per)
}