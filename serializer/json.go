package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJson(message proto.Message) (json string, err error) {
	marshaller := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent:       "	",
	}

	return marshaller.MarshalToString(message)
}
