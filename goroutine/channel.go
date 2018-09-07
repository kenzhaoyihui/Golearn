package main

import (
	"fmt"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}

func add(ch chan int, x, y int) {
	z := x + y
	ch <- z
	fmt.Println(z)
}


func main() {
	//chs := make([]chan int, 10)
	//for i := 0; i < 10; i++ {
	//	chs[i] = make(chan int)
	//	go add(chs[i], i ,i)
	//}
	//
	//for _, ch := range(chs) {
	//	<- ch
	//}

	ch := make(chan int, 10)
	ch <- 1


	//select {
	//case ch <- 0:
	//	fmt.Println("0 channel")
	//case ch <- 1:
	//	fmt.Println("1 channel")
	//}
	//i := <- ch
	//fmt.Println("Value received: ", i)

	//timeout := make(chan bool, 1)
	//go func() {
	//	time.Sleep(10e9)
	//	timeout <- true
	//}()
	//
	//select {
	//	case <- ch:
	//		fmt.Println("receive the data from channel ch ")
	//	case <- timeout:
	//		fmt.Println("Timeout!")
	//}

	ch <- 2
	close(ch)

	for x := range ch {
		fmt.Println(x)
	}

	if x, ok := <- ch; !ok {
		fmt.Println("Closed!")
	}else {
		fmt.Println(x, ok)
	}


}
