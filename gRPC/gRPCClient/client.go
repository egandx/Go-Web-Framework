package main

import (
	"client.go/services"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const address = "127.0.0.1:8082"

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
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
