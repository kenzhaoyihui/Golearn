package main

import (
	"net/rpc"
	"net"
	"log"
	"my/rpc/newrpc"
)

//const HelloServiceName = "pkg.HelloService"
//
//type HelloServiceInterface interface {
//	Hello(request string, reply *string) error
//}
//
//type HelloService1 struct {}
//
//func (p *HelloService1) Hello(request string, reply *string) error{
//	*reply = "hello: " + request
//	return nil
//}
//
//
//func RegisterHelloService(svc HelloServiceInterface) error {
//	return rpc.RegisterName(HelloServiceName, svc)
//}

func main(){
	//rpc.RegisterName("HelloService", new(HelloService))

	newrpc.RegisterHelloService(new(newrpc.HelloService1))


	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go rpc.ServeConn(conn)
	}
}
