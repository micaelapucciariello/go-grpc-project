package service

import (
	"context"
	"github.com/micaelapucciariello/grpc-project/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PCServer struct {
}

func NewPCServer() *PCServer {
	return &PCServer{}
}

func (s *PCServer) CreatePC(context.Context, *pb.CreatePCRequest) (*pb.CreatePCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePC not implemented")
}
