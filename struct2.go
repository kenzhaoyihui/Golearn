package main

import "fmt"

type human struct{
	Sex int
	Gg string
}

type teacher struct{
	human
	Name string
	Age int
}

type student struct{
	human
	Name string
	Age int
}

func main(){
	a := teacher{Name: "zyh", Age:22, human: human{Sex: 0, Gg: "love"}}
	b := student{Name: "lj", Age:23, human: human{Sex: 1, Gg: "you"}}
	a.Name = "aaa"
	a.Age = 12
	a.Sex = 100
	a.Gg = "loveyou"
	fmt.Println(a, b)
}