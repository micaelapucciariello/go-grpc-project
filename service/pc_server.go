package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/micaelapucciariello/grpc-project/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type PCServer struct {
	Store PCStore
}

func NewPCServer() *PCServer {
	return &PCServer{}
}

func (server *PCServer) CreatePC(ctx context.Context, req *pb.CreatePCRequest) (*pb.CreatePCResponse, error) {
	pc := req.GetPc()
	log.Printf("create pc request with id: %v", pc.Id)

	if len(pc.Id) > 0 {
		_, err := uuid.Parse(pc.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid pc uuid: %s", pc.Id)
		} else {
			pc.Id = uuid.New().String()
		}
	}

	err := server.Store.Save(pc)
	if err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "item already exists")
		}
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	res := &pb.CreatePCResponse{Id: pc.Id}

	return res, nil
}
