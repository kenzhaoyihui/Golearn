package main

import (
	"os"
	"fmt"
	"flag"
	"strconv"
)

var goal int


func primetask(c chan int) {
	p := <- c
	if p > goal {
		os.Exit(0)
	}

	fmt.Println(p)

	nc := make(chan int)
	go primetask(nc)

	for {
		i := <- c
		if i%p !=0 {
			nc <- i
		}
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	if args != nil && len(args) > 0 {
		_, err := strconv.Atoi(args[0])

		if err != nil {
			goal = 100
		}
	}else {
		goal = 100
	}

	fmt.Println("goal=" , goal)

	c := make(chan int)
	go primetask(c)

	for i:=2; ; i++ {
		c <- i
	}
}
//package main
//
//import "fmt"
//
//type Describer interface {
//	Describe()
//}
//type St string
//
//func (s St) Describe() {
//	fmt.Println("被调用le!")
//}
//
//func findType(i interface{}) {
//	switch v := i.(type) {
//	case Describer:
//		v.Describe()
//	case string:
//		fmt.Println("String 变量")
//	default:
//		fmt.Printf("unknown type\n")
//	}
//}
//
//func main() {
//	findType("Naveen")
//	st := St("我的字符串")
//	findType(st)
//}