package main

import (
	"context"
	"flag"
	"github.com/micaelapucciariello/grpc-project/pb"
	"github.com/micaelapucciariello/grpc-project/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pc := sample.NewPC()
	req := &pb.CreatePCRequest{Pc: pc}

	res, err := pcClient.CreatePC(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Fatal("pc already exists")
		} else {
			log.Fatal("cannot create pc")
		}

		return
	}

	log.Printf("created pc with id: %v", res.Id)

}
