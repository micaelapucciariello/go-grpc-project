package serializer_test

import (
	"github.com/micaelapucciariello/grpc-project/pb"
	"github.com/micaelapucciariello/grpc-project/sample"
	"github.com/micaelapucciariello/grpc-project/serializer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSerializeMessage(t *testing.T) {
	t.Parallel()

	laptop1 := sample.NewPC()
	binFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	err := serializer.WriteProtobufToBinaryFile(laptop1, binFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, binFile)

	laptop2 := &pb.PC{}
	err = serializer.ReadProtobufToBinaryFile(laptop2, binFile)
	assert.NoError(t, err)
	assert.Equal(t, laptop1.Id, laptop2.Id)

	err = serializer.WriteProtobufToJsonFile(laptop1, jsonFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonFile)
}
