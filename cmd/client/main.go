package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	service_v1 "grpc-exercises/service/v1"
	"log"
	"os"
)

func main() {
	serverUrl := os.Getenv("SERVER_URL")
	if serverUrl == "" {
		log.Fatalf("SERVER_URL is required")
	}

	fmt.Printf("Connecting client to %s\n", serverUrl)
	clientConnection, err := grpc.Dial(serverUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s with error: %v", serverUrl, err)
	}

	defer clientConnection.Close()

	client := service_v1.NewSumServiceClient(clientConnection)

	request := &service_v1.SumRequest{
		First: 1,
		Second: 2,
	}

	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling server: %v", err)
	}

	log.Printf("1 + 2 = %d", response.Sum)
}