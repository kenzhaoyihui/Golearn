// map.go
package main

import "fmt"

func main(){
	//var m map[int]string = make(map[int]string)
	//m = map[int]string{}
	/*
	m := make(map[int]string)
	m[1] = "OK"
	delete(m, 1)
	a := m[1]
	fmt.Println(a)
	*/

	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	//m[1] = make(map[int]string)
	//m[1][1] = "OK"
	//a := m[1][1]
	a, ok := m[2][1]
	if !ok{
		m[2] = make(map[int]string)
	}
	m[2][1] = "good"
	a, ok = m[2][1]
	fmt.Println(a, ok)
}