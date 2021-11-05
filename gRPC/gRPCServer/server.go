package main

import (
	"gRPCServer/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

const port = ":8082"

func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	grpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(grpcServer,&services.ProdService{})
	log.Println("start gRPC listen on port " + port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}


}
