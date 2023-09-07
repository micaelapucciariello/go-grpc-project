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

func New(s PCStore) *PCServer {
	return &PCServer{
		Store: s,
	}
}

func (server *PCServer) CreatePC(ctx context.Context, req *pb.CreatePCRequest) (*pb.CreatePCResponse, error) {
	pc := req.GetPc()
	log.Printf("create pc request with id: %v", pc.Id)

	if len(pc.Id) == 0 {
		pc.Id = uuid.New().String()
	}

	if _, err := uuid.Parse(pc.Id); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pc uuid: %s", pc.Id)
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

func (server *PCServer) SearchPC(req *pb.SearchPCRequest, stream pb.PCService_SearchPCServer) error {
	filter := req.GetFilter()
	err := server.Store.Search(filter, func(pc *pb.PC) error {
		res := &pb.SearchPCResponse{Pc: pc}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		log.Printf("sent pc with id: %v", pc.Id)
		return nil
	},
	)
	if err != nil {
		return err
	}

	return nil
}

// mustEmbedUnimplementedPCServiceServer implements pb.PCServiceServer.
func (server *PCServer) mustEmbedUnimplementedPCServiceServer() {
	panic("unimplemented")
}
