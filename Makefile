gen:
	protoc --proto_path=./proto ./proto/*.proto --plugin=$(go env GOPATH)/bin/protoc-gen-go --go_out=./

clean:
	rm pb/*.go

test:
	go test -v ./...

run:
	go run main.go


.PHONY: gen, clean, run, test