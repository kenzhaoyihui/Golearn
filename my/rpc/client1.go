package main

import (
	//"net/rpc"
	//"github.com/labstack/gommon/log"
	"fmt"
	"log"
	"my/rpc/newrpc"
)

//type HelloServiceClient struct {
//	*rpc.Client
//}
//
//var _ HelloServiceInterface = (*HelloServiceClient)(nil)
//
//
//
//func DialHelloService(network, address string) (*HelloServiceClient, error) {
//	c, err := rpc.Dial(network, address)
//	if err != nil {
//		return nil, err
//	}
//
//	return &HelloServiceClient{Client: c}, nil
//}
//
//func (p *HelloServiceClient) Hello(request string, reply *string) error {
//	return p.Client.Call("pkg.HelloService.Hello", request, &reply)
//}


func main() {

	client, err := newrpc.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	var reply string
	//err = client.Call( HelloServiceName + ".Hello", "I am yzhao", &reply)
	err = client.Hello(  "I am yzhao123", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}