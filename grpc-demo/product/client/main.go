package main

import (
	"context"
	"product/product"

	"log"

	"google.golang.org/grpc"
)

const address = "127.0.0.1:9527"

func AddProduct(ctx context.Context, client product.ProductInfoClient) (id string) {
	aMac := &product.Product{Name: "MacBookPro 2021", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, aMac)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success,id=", productId.Value)
	return productId.Value
}

func GetProduct(ctx context.Context, client product.ProductInfoClient, id string) {
	p, err := client.GetProduct(ctx, &product.ProductId{Value: id})
	if err != nil {
		log.Println("get product,err.", err)
		return
	}
	log.Printf("get product success: %+v\n", p)
}

func main(){
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("dial connect fail.", err)
		return
	}
	client := product.NewProductInfoClient(conn)
	ctx := context.Background()

	id := AddProduct(ctx, client)
	GetProduct(ctx, client, id)
}

