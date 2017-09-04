package main

import (
	"fmt"
	//"time"
	"runtime"
)


func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool, 10)
	for i:=0; i<10; i++{
		go Go(c, i)
	}

	//设置一个缓存长度为10的channel 
	for i:=0;i<10;i++{
		<-c
	}
	
}

func Go(c chan bool, index int){
	a := 1
	for i:=0; i<100000000; i++{
		a += i
	}
	fmt.Println(index, a)

	c <- true
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