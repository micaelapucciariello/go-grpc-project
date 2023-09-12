package main

import (
	"context"
	"flag"
	"github.com/micaelapucciariello/grpc-project/pb"
	"github.com/micaelapucciariello/grpc-project/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func createPC(pcClient pb.PCServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pc := sample.NewPC()
	req := &pb.CreatePCRequest{Pc: pc}

	res, err := pcClient.CreatePC(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Fatal("pc already exists")
		} else if !ok {
			log.Fatal("cannot create pc")
		}

		return
	}

	log.Printf("created pc with id: %v", res.Id)
}

func searchPC(pcClient pb.PCServiceClient, filter *pb.Filter) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchPCRequest{
		Filter: filter,
	}

	stream, err := pcClient.SearchPC(ctx, req)
	if err != nil {
		log.Fatalf("cannot search pc: %v", err.Error())
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("cannot receive response: %v", err)
			return
		}

		pc := res.GetPc()
		log.Printf("- found pc - id: %v", pc.GetId())
		log.Printf("+ price: %v", pc.GetUsdPrice())
	}
}

func main() {
	serverAddress := flag.String("address", "0.0.0.0:8080", "server address")
	if *serverAddress == "" || &serverAddress == nil {
		log.Fatal("invalid address")
		return
	}
	flag.Parse()
	log.Printf("dial server on server address: %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server")
	}

	pcClient := pb.NewPCServiceClient(conn)
	for i := 0; i < 10; i++ {
		createPC(pcClient)
	}

	filter := &pb.Filter{
		MaxPriceUsd: 1500,
		MinCpuCores: 6,
		MinCpuGhz:   3,
	}

	searchPC(pcClient, filter)
}
