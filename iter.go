package main

import "fmt"
//import "sort"

func main(){
	//for k, v :=range map{}
	//sm := make([]map[int]string, 5)
	/*
	for _,v := range sm{
		v = make(map[int]string, 1)
		v[1] = "OK"
		fmt.Println(v)
	}
	fmt.Println(sm)
	*/

	/*
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
	*/

	/*
	m := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	s := make([]int, len(m), 100)
	s1 := make([]string, len(m), 100)
	i := 0
	for k,v := range m{
		s[i] = k
		s1[i] = v
		i++
	}
	sort.Ints(s)
	sort.Strings(s1)
	fmt.Println(s)
	fmt.Println(s1)
	*/

	m := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	m1 := make(map[string]int)
	for k, v := range m{
		m1[v] = k
	}
	fmt.Println(m1)
}