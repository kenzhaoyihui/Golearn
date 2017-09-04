package main

import (
	"fmt"
	//"time"
	"runtime"
	"sync"
)


func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c := make(chan bool, 10)
	//利用sync包的WaitGroup方法
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i:=0; i<10; i++{
		go Go(&wg, i)
	}
	wg.Wait()
	//for i:=0;i<10;i++{
	//	<-c
	//}
	
}

func Go(wg *sync.WaitGroup, index int){
//func Go(c chan bool, index int){
	a := 1
	for i:=0; i<100000000; i++{
		a += i
	}
	fmt.Println(index, a)

	//c <- true
	wg.Done()
}
/*
func main(){
	c := make(chan bool)
	go func(){
		fmt.Println("Go Go Go!!!")
		c <- true
		close(c)
	}()

	for v := range c{
		fmt.Println(v)
	}
	//<-c
	//time.Sleep(2 * time.Second)

}
*/