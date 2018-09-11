package newrpc

import (
	"net/rpc"
	//"github.com/labstack/gommon/log"
	"fmt"
)

type HelloServiceClient struct {
	CC *rpc.Client
}

/**

type T struct{}
var _ I = T{}
其中 I为interface。上面用来判断 type T是否实现了I,用作类型断言，如果T没有实现借口I，则编译错误
 */

var _ HelloServiceInterface = (*HelloServiceClient)(nil)



func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{CC: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	//return p.Client.Call("pkg.HelloService.Hello", request, &reply)
	return p.CC.Call(HelloServiceName + ".Hello", request, &reply)
}

//type T struct {
//
//}
//
//type I interface {
//	print(name string)
//}
//
////var _ I = (*T)(nil)
//
//var _ I = new(T) // === var _ I = (*T)(nil)
//
//func (p *T) print(name string)  {
//	fmt.Println(name)
//}