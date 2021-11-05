package main

import (
	"client.go/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

const address = "127.0.0.1:8082"

func main() {

	creds, err := credentials.NewClientTLSFromFile("keys/grpcserver.crt","egan.com")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Println("dial connect fail.", err)
		return
	}

	defer conn.Close()

	client := services.NewProdServiceClient(conn)

	stock, err := client.GetProdStock(context.Background(),&services.ProdRequest{ProdId: 1} )
	if err != nil {
		return 
	}

	fmt.Println(stock.ProdStock)
}
