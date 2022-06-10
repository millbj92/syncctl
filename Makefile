BINARY_NAME=synctl

build_server:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux_amd64 ./cmd/server/main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin_amd64 ./cmd/server/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows_amd64 ./cmd/server/main.go

build_cli:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-cli-linux_amd64 ./cmd/cli/main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME}-cli-darwin_amd64 ./cmd/cli/main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-cli-windows_amd64 ./cmd/cli/main.go

run:
	./{BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-linux_amd64
	rm ${BINARY_NAME}-darwin_amd64
	rm ${BINARY_NAME}-windows_amd64
