gen:
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=./ --go-grpc_out=.

clean:
	rm pb/*.go

test:
	go test -v -cover -race ./...

server:
	go run ./cmd/server/main.go port 8080

client:
	go run ./cmd/client/main.go address 0.0.0.0:8080


.PHONY: gen, clean, server, client, test