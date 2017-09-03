package main

import "fmt"

func main(){
	LABLE1:
		for {
			for i:=0; i<10; i++{
				for{
					fmt.Println(i)
					continue LABLE1
				}
			}
		}
		//LABLE1:
		//	fmt.Println("OK")
}