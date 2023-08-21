package server

import (
	"flag"
	"fmt"
	"github.com/micaelapucciariello/grpc-project/pb"
	"github.com/micaelapucciariello/grpc-project/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 0, "port value")
	flag.Parse()
	log.Printf("start server on port: %d", port)

	pcServer := service.New(service.NewInMemoryPCStore())
	grpcServer := grpc.NewServer()
	pb.RegisterPCServiceServer(grpcServer, pcServer)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server")
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
