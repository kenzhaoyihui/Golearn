package newrpc

import (
	"net/rpc"
)


const HelloServiceName = "pkg.HelloService"

type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

type HelloService1 struct {}

func (p *HelloService1) Hello(request string, reply *string) error{
	*reply = "hello: " + request
	return nil
}


func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}
