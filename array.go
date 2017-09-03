package main

import "fmt"

const MAX int = 3
func main(){
	a := []int{10,20,30}
	var p [MAX]*int

	for i:=0;i<MAX;i++{
		p[i] = &a[i]
	}

	for i:=0;i<MAX;i++{
		fmt.Printf("a[%d] = %d\n", i, *p[i])
	}

	/*
	a := [...]int{3,7,1,9,3,5}
	fmt.Println(a)

	num := len(a)
	for i := 0; i<num; i++{
		for j := i+1; j<num; j++{
			if a[i]>a[j]{
				temp := a[i]
				a[i] =a[j]
				a[j] = temp
			} 
		}
	}
	fmt.Println(a)
	*/
}