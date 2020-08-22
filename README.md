# GRPC Exercises

My attempt for the exercises from [Udemy gRPC Golang course](https://www.udemy.com/course/grpc-golang/learn/lecture/11018796#overview)

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
