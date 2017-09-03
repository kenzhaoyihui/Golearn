package main

import "fmt"
import "reflect"

type User struct{
	Id int
	Name string
	Age int
}

type Manager struct{
	User
	title string
}

func main(){
	m := Manager{User:User{1,"OK",2}, title: "helloworld"}
	t := reflect.TypeOf(m)

	//fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0,1}))

}