package main

import (
	"reflect"
	"fmt"
)

type Enum int

const (
	Zero Enum = 0
)

type Cat struct {
	Name string
	//Type int `json: "yzhao" id:"100"`  // must not have the blank after ":"
	Type int `json:"yzhao" id:"100"`  // this is right
}


func add(a, b int) int {
	return a + b
}

func main() {
	var a int

	typeOfA := reflect.TypeOf(a)

	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	typeOfCat := reflect.TypeOf(Cat{
		Name:"lj",
	})

	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	typeOfEnum := reflect.TypeOf(Zero)
	fmt.Println(typeOfEnum.Name(), typeOfEnum.Kind())

	fmt.Println("----------------------------------")

	ins := &Cat{Name:"lljj"}
	typeIns := reflect.TypeOf(ins)
	fmt.Printf("name: %v, kind: %v", typeIns.Name(), typeIns.Kind())

	fmt.Println()
	typeIns = typeIns.Elem()
	fmt.Printf("name: %v, kind: %v", typeIns.Name(), typeIns.Kind())


	fmt.Println()
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		fmt.Println(catType.Tag.Get("json"))
	}


	fmt.Println()

	valueZero := reflect.ValueOf(Zero)

	getZero := valueZero.Interface().(Enum)  // should be type , not kind

	getZero1 := int(valueZero.Int())

	fmt.Println(valueZero, getZero, getZero1)


	a = 1024
	valueA := reflect.ValueOf(&a)
	valueA.Elem().SetInt(24)
	fmt.Println(a)
	//valueA = valueA.Elem()
	//valueA.SetInt(245)
	//fmt.Println(valueA.Int())

	valueAA := reflect.TypeOf(a)
	aIns := reflect.New(valueAA)
	fmt.Println(aIns.Type(), aIns.Kind())


	// use reflect to invoke the method add()
	funcAdd := reflect.ValueOf(add)
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

	retList := funcAdd.Call(paramList)
	fmt.Println(retList)

	for index, result := range retList {
		fmt.Println(index, result)
	}
}
