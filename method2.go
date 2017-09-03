package main

import "fmt"

type TZ int

func main(){
	var a TZ
	a.Incress(100)
	fmt.Println(a)
}

func (tz *TZ)Incress(num int){
	*tz += TZ(num)
}