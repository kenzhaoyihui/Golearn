package main

import "fmt"

func main(){
	a := []int{1,2,3,4,5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s2 = append(s2, 1,2,2,3,4,5,6,7)
	s1[0] = 9
	fmt.Println(s1, s2)

	s3 := []int{1,2,3,4,5,6}
	s4 := []int{7,8,9}
	copy(s4, s3)
	fmt.Println(s4)
	/*
	s1 := make([]int, 3, 6)
	fmt.Printf("%p", s1)
	s1 = append(s1, 1,2,3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	*/
	//var s1 []int
	//fmt.Println(s1)
	//a := [10]int{1,2,3,4,5,6,7,8,9,10}
	//fmt.Println(a)
	//s1 := a[:5]
	//fmt.Println(s1)

	//s1 := make([]int, 3, 10)
	//fmt.Println(len(s1), cap(s1))
}