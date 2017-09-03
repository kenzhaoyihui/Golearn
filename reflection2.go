package main

import "fmt"
import "reflect"


type User struct{
	Id int
	Name string
	Age int
}

func Set(o interface{}){
	v := reflect.ValueOf(o)

	if v.Kind()==reflect.Ptr&& !v.Elem().CanSet(){
		fmt.Println("XXX")
		return
	}else{
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid(){
		fmt.Println("BAD")
		return
	}

	if f.Kind()==reflect.String{
		f.SetString("ByeBye")
	}
}

func main(){
	u := User{1, "ok", 12}
	Set(&u)
	fmt.Println(u)
	/*
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x)
	*/
}