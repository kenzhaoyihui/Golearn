package main

import (
	"net/rpc"
	//"github.com/labstack/gommon/log"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {

	//client, err := rpc.Dial("tcp", "localhost:1234")
	conn, err := net.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("dialing: ", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("HelloService.Hello", "I am yzhao", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}