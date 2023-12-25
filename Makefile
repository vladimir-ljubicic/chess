BINARY_NAME=chess

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin main.go
run: build
	./bin/${BINARY_NAME}-darwin
clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin