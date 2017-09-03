package main

import "fmt"

const a int = 1
const b = 'A'

const (
	_MAX_COUNT = 12345
	cZYH = 34567
	//_ and c 不让外面访问，之让内部包的使用
	c = "123"
	d = len(c)
	e, f = "lj", 23
	g, h
	i = 2
	j = iota
	k = iota
	l = iota
	m = iota
	n = iota
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(!true)
	fmt.Println(6&11)
	fmt.Println(6|11)
	fmt.Println(6^11)
	fmt.Println(6&^11)
	x := 1
	if x>0 && 10/x>1{
		fmt.Println("OK")
	}
}