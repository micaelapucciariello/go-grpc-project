package service_test

import (
	"context"
	"github.com/micaelapucciariello/grpc-project/pb"
	"github.com/micaelapucciariello/grpc-project/sample"
	"github.com/micaelapucciariello/grpc-project/serializer"
	"github.com/micaelapucciariello/grpc-project/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func CreateClientServerPC(t *testing.T) {
	t.Parallel()

	server, address := StartTestPCServer(t)
	client := newTestPCClient(t, address)

	// test client
	pc := sample.NewPC()
	expectedID := pc.GetId()
	req := &pb.CreatePCRequest{
		Pc: pc,
	}

	res, err := client.CreatePC(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// find item
	other, err := server.Store.Find(pc.Id)
	require.NoError(t, err)
	require.NotNil(t, other)
	require.Equal(t, expectedID, other.Id)

	// require created item and found item are the same
	requireSamePC(t, pc, other)

}

func newTestPCClient(t *testing.T, address string) pb.PCServiceClient {
	conn, err := grpc.Dial(address)
	require.NoError(t, err)

	return pb.NewPCServiceClient(conn)
}

func StartTestPCServer(t *testing.T) (*service.PCServer, string) {
	server := service.New(service.NewInMemoryPCStore())

	grpcServer := grpc.NewServer()
	pb.RegisterPCServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", "0")
	require.NoError(t, err)

	go func() {
		err := grpcServer.Serve(listener)
		require.NoError(t, err)
	}()

	return server, listener.Addr().String()
}

func requireSamePC(t *testing.T, pc1 *pb.PC, pc2 *pb.PC) {
	json1, err := serializer.ProtobufToJson(pc1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJson(pc2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
