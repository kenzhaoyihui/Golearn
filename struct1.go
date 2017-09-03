package main

import "fmt"


type person struct{
	Name string
	Age int
	Contact struct{
		Phone, City string
	}
}
func main(){
	a := person{Name: "zyh", Age:22}
	a.Contact.Phone = "18021521615"
	a.Contact.City = "Nanjing"
	fmt.Println(a)
	/*
	a := &struct{
		Name string
		Age int
	}{
		Name: "zyh",
		Age: 22,
	}
	fmt.Println(a)
	*/
}