package main

import (
	"sync"
	"fmt"
)

var a string

var once sync.Once

func setup() {
	a = "yzhao"
	fmt.Println("once:", a)
}



func doPrint(ch chan int) {
	ch <- 1
	once.Do(setup)
	println(a)
}

func twoPrint() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go doPrint(chs[i])
	}

	for _, ch := range(chs) {
		<- ch
	}
}

func main() {
	twoPrint()
}
