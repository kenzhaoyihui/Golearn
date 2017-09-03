package main

import "fmt"
import "reflect"

type User struct{
	Id int
	Name string
	Age int
}

func (u User)Hello(name string, age int){
	fmt.Println("Hello", name, ",my name is:" , u.Name, "...Age:", age ,"my age is :",u.Age)
}

func main(){
	u := User{1, "hyzsherry", 23}
	//u.Hello("Zyh")
	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := make([]reflect.Value, 2)
	args[0] = reflect.ValueOf("joe")
	args[1] = reflect.ValueOf(29)
	mv.Call(args)

	//args := []reflect.Value{reflect.ValueOf("joe")}
	//mv.Call(args)

}