package service_test

import (
	"context"
	"github.com/micaelapucciariello/grpc-project/pb"
	"github.com/micaelapucciariello/grpc-project/sample"
	"github.com/micaelapucciariello/grpc-project/serializer"
	"github.com/micaelapucciariello/grpc-project/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"io"
	"net"
	"testing"
)

func CreateClientServerPC(t *testing.T) {
	t.Parallel()

	store := service.NewInMemoryPCStore()
	server, address := StartTestPCServer(t, store)
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

func StartTestPCServer(t *testing.T, store service.PCStore) (*service.PCServer, string) {
	server := service.New(store)

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

func TestClientSearchPC(t *testing.T) {
	t.Parallel()

	store := service.NewInMemoryPCStore()
	expectedIDs := make(map[string]bool)

	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Value: "8",
			Unit:  pb.Memory_GBYTE,
		},
	}

	for i := 0; i < 6; i++ {
		pc := sample.NewPC()

		switch i {
		case 0:
			pc.UsdPrice = 2500
		case 1:
			pc.Cpu.Cores = 2
		case 2:
			pc.Cpu.MinGhz = 2
		case 3:
			pc.Memory[0] = &pb.Memory{
				Value: "4",
				Unit:  pb.Memory_MBYTE,
			}
		case 4:
			pc.UsdPrice = 1000
			pc.Cpu.Cores = 6
			pc.Cpu.MinGhz = 2.5
			pc.Cpu.MaxGhz = 3
			pc.Memory[0] = &pb.Memory{
				Value: "16",
				Unit:  pb.Memory_GBYTE,
			}
			expectedIDs[pc.Id] = true
		case 5:
			pc.UsdPrice = 1300
			pc.Cpu.Cores = 4
			pc.Cpu.MinGhz = 3
			pc.Cpu.MaxGhz = 3.5
			pc.Memory[0] = &pb.Memory{
				Value: "64",
				Unit:  pb.Memory_GBYTE,
			}
			expectedIDs[pc.Id] = true
		}

		err := store.Save(pc)
		require.NoError(t, err)
	}

	_, address := StartTestPCServer(t, store)
	pcClient := newTestPCClient(t, address)

	req := &pb.SearchPCRequest{Filter: filter}
	stream, err := pcClient.SearchPC(context.Background(), req)
	require.NoError(t, err)

	found := 0

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetPc().GetId())
		found += 1
	}

	require.Equal(t, len(expectedIDs), found)
}

func requireSamePC(t *testing.T, pc1 *pb.PC, pc2 *pb.PC) {
	json1, err := serializer.ProtobufToJson(pc1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJson(pc2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}
