package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

type HelloService struct {}

func (p *HelloService) Hello(request string, reply *string) error{
	*reply = "hello: " + request
	return nil
}


func main(){

	//srv := rpc.NewServer()
	////rpc.RegisterName("HelloService", new(HelloService))
	//srv.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp error: ", err)
	}
	defer listener.Close()

	//srv := rpc.NewServer()
	rpc.RegisterName("HelloService", new(HelloService))
	//srv.RegisterName("HelloService", new(HelloService))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		//go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}