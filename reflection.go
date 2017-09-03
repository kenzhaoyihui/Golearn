package main

import "fmt"
import "reflect"

type User struct{
	Id int
	Name string
	Age int
}

func (u User) Hello(){
	fmt.Println("Hello world!")
}

func Info(o interface{}){
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	if k:= t.Kind(); k!=reflect.Struct{
		fmt.Println("XX")
		return
	}

	v := reflect.ValueOf(o)
	fmt.Println("Field:")

	for i:=0;i < t.NumField();i++{
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i:=0;i<t.NumMethod();i++{
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
func main(){
	u := User{1, "OK", 12}
	Info(&u)
}