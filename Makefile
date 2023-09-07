gen:
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=./ --go-grpc_out=.

clean:
	rm pb/*.go

test:
	go test -v -cover -race ./...

server:
	go run /cmd/server/main.go

client:
	go run /cmd/client/main.go


.PHONY: gen, clean, server, client, test