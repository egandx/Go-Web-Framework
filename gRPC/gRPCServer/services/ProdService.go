package services

import (
	"context"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ProdService struct {

}

func (s *ProdService) mustEmbedUnimplementedProdServiceServer() {
	panic("implement me")
}

func (s *ProdService) GetProdStock(ctx context.Context,request *ProdRequest) (*ProdResponse, error)  {

	return &ProdResponse{
		state:         protoimpl.MessageState{},
		sizeCache:     0,
		unknownFields: nil,
		ProdStock:     255 ,
	}, nil
}
