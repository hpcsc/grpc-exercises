# GRPC Exercises

My attempt for the exercises from [Udemy gRPC Golang course](https://www.udemy.com/course/grpc-golang)

- GRPC proto files are at: `./service/v1/*.proto`
- Generated files are placed in the same location with proto files
- Same server implements all service interfaces: `./cmd/server/main.go`
- Multiple clients connect to the same server: `./cmd/client/main.go`
- Current services:
  - Sum: Unary GRPC Service
  - Prime Number Decomposition: Server Streaming
  - Compute Average: Client Streaming
  - Find Maximum: Bi-Directional Streaming

![alt Client UI](https://github.com/hpcsc/grpc-exercises/raw/master/client-ui.png)

## Generate GRPC code

```
./batect generate
```

## Run GRPC server

```
./batect startServer
```

This task will invoke `generate` task before running the server

## Run GRPC client

```
./batect startClient
```

This task will invoke `generate` task and run server before starting client
