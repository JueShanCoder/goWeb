package main

import (
	base1 "geerpc"
	"log"
	"net"
)

func main() {
	addr := make(chan string)
	go startServer(addr)
}

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}

	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	base1.Accept(l)
}
