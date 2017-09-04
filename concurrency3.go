package main

import "fmt"
import "time"


func main(){
	c := make(chan bool)
	select{
		case v := <-c:
			fmt.Println(v)
		case <-time.After(3*time.Second):
			fmt.Println("Timeout")
	}
}
/*
func main(){
	c := make(chan int)
	go func(){
		for v := range c{
			fmt.Println(v)
		}
	}()

	for{
		select{
			case c<-0:
			case c<-1:
		}
	}
}
*/