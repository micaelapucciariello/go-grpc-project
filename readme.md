# grpc-project
## pcbook

### Setup local environment

_Setup in macOS with Homebrew._

#### protoc
run `go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`

#### proto-gen
run `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` `export PATH="$PATH:$(go env GOPATH)/bin"`

#### linter
run `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

#### pre-commit
run `pip install pre-commit` `pre-commit install`