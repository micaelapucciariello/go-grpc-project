package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message: %s", err)
	}

	err = ioutil.WriteFile(filename, data, 0664)
	if err != nil {
		return fmt.Errorf("cannot write data into file: %s", err)
	}

	return nil
}

func WriteProtobufToJsonFile(message proto.Message, filename string) error {
	data, err := ProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("cannot convert protobuf to json: %s", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0664)
	if err != nil {
		return fmt.Errorf("cannot write data into file: %s", err)
	}

	return nil
}

func ReadProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read data from file: %s", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal data into proto message: %s", err)
	}

	return nil
}
