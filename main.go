//go:generate protoc --go_out=plugins=grpc:./services/user user.proto

package main

import (
	"google.golang.org/grpc"
	"grpc-test/services"
	"log"
	"net"
)

func main() {
	tcpServer, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error in starting TCP server. Error:", err.Error())
		return
	}
	gServer := grpc.NewServer()
	services.NewServices(gServer)
	if err := gServer.Serve(tcpServer); err != nil {
		log.Fatal("Error in serve gRPC in TCP Server. Error:", err.Error())
	}
}


