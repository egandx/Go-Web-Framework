package main

import (
	"gRPCServer/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

const port = ":8082"

func main() {

	creds, err := credentials.NewServerTLSFromFile("keys/grpcserver.crt","keys/grpcserver_no_password.key")
	if err != nil {
		log.Fatalln(err)
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(grpcServer,&services.ProdService{})
	log.Println("start gRPC listen on port " + port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}


}
