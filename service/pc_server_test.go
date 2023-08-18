package service

import (
	"context"
	"github.com/micaelapucciariello/grpc-project/pb"
	"github.com/micaelapucciariello/grpc-project/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestPCServer_CreatePC(t *testing.T) {
	t.Parallel()

	PCWithoutID := sample.NewPC()
	PCWithoutID.Id = ""

	PCWithInvalidID := sample.NewPC()
	PCWithInvalidID.Id = "no_valid_id"

	duplicatedPC := sample.NewPC()
	duplicatedStore := NewInMemoryPCStore()
	err := duplicatedStore.Save(duplicatedPC)
	assert.NoError(t, err)

	testCases := []struct {
		name  string
		pc    *pb.PC
		store *InMemoryPCStore
		code  codes.Code
	}{
		{
			name:  "success_with_id",
			pc:    sample.NewPC(),
			store: NewInMemoryPCStore(),
			code:  codes.OK,
		},
		{
			name:  "success_without_id",
			pc:    PCWithoutID,
			store: NewInMemoryPCStore(),
			code:  codes.OK,
		},
		{
			name:  "error_invalid_id",
			pc:    PCWithInvalidID,
			store: NewInMemoryPCStore(),
			code:  codes.InvalidArgument,
		},
		{
			name:  "error_already_exists",
			pc:    duplicatedPC,
			store: duplicatedStore,
			code:  codes.AlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		ctx := context.Background()

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreatePCRequest{Pc: tc.pc}

			server := NewPCServer(tc.store)
			pc, err := server.CreatePC(ctx, req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotEmpty(t, pc)
				require.Equal(t, tc.pc.Id, pc.Id)
			} else {
				require.Error(t, err)
				require.Nil(t, pc)
				st, ok := status.FromError(err)
				if ok {
					require.Equal(t, st.Code().String(), tc.code.String())
				}
			}

		},
		)
	}
}
