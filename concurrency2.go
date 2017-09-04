package main

import (
	"fmt"
	//"runtime"
	//"sync"
)

func main(){
	o := make(chan bool, 2)
	c1,c2 := make(chan int), make(chan string)
	go func(){
		//a,b := false,false
		for{
			select{
				case v,ok:=<-c1:
					if !ok{
						//if !a{
							//a = true
						o<-true
						//}
						//fmt.Println("c1")
						//o<- true
						//a = true
						break
					}
					fmt.Println("c1", v)
				case v,ok:=<-c2:
					if !ok{
						//if !b{
							//b = true
						o<-true
						//}
						//fmt.Println("c2")
						//o<- true
						//b = true
						break
					}
					fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	//close(c2)

	for i:=0; i<2; i++{
		<-o
	}

}