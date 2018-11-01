package main

import (
	"fmt"
	"net/rpc"
	"net"
	"net/rpc/jsonrpc"
)


func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil{
		panic(err)
	}
	return jsonrpc.NewClient(conn), nil
}

func main() {
	var a map[int]string
	b := make(map[string]string)
	c := []byte{'a','b'}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
