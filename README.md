# GRPC Sum API

A test GRPC application

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
