# Go gRPC Microservice: Product Service

This is a simple microservice written in Go using gRPC. It exposes a `GetProduct` method to fetch product details by ID.

## Features

- gRPC server implementation in Go
- Simple in-memory product store
- Client call to demonstrate gRPC functionality


## Prerequisites

- Go 1.18+

## Protocol Buffers compiler Prerequisites
To generate the Go code from your .proto files and place them in the productService/productpb directory, run the following commands from the root of your project:

- [Protocol Buffers compiler (`protoc`)](https://grpc.io/docs/protoc-installation/)
- Go plugins for protoc:
```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## How to Run
1. Clone the repository and navigate into productService:
```bash
  git clone https://github.com/yourusername/go-microservice.git
  cd go-microservice/productService
```

2. Run the application with Makefile:
```bash
  make run
```
- Or with go command

```bash
  go run main.go
```

## This will:

- Start the gRPC server on port 4001

- Call the GetProduct method using a client request after a short delay