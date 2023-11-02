export GOPROXY="https://proxy.golang.org,direct"
export GOPRIVATE="github.com/ioriver"
export GONOSUMDB="github.com/ioriver/*"
go mod init github.com/ioriver/ioriver-go
go mod tidy
go mod vendor
go build ./...
go test ./...