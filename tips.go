package main

import (
	"fmt"
	//"time"
)

func main(){
	s := []string{"a", "b", "c"}
	for _,v:=range s{
		go func(v string){
			fmt.Println(v)
		}(v)
	}
	select{}
}

/*
func main(){
	t := time.Now()
	fmt.Println(t.Format(time.ANSIC))
}
*/


/*
func Pingpong(s []int) []int {
	s = append(s, 3,4,5)
	return s
}

func main(){
	s := make([]int,0)
	fmt.Println(s)
	s = Pingpong(s)
	fmt.Println(s)
}
*/