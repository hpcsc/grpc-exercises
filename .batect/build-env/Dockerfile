FROM golang:1.14.7-stretch

RUN apt-get -y update && \
  apt-get install -y --no-install-recommends protobuf-compiler=3.0.0-9 && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

ENV GO111MODULE=on

RUN go get google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0 && \
  go get google.golang.org/grpc/cmd/protoc-gen-go-grpc@2983360 && \
  mv /go/bin/* /usr/bin

WORKDIR /app
