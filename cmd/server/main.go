package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	service_v1 "grpc-exercises/service/v1"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

type server struct {
	service_v1.UnimplementedSumServiceServer
	service_v1.UnimplementedComputeAverageServiceServer
}

func (s *server)Sum(ctx context.Context, req *service_v1.SumRequest) (*service_v1.SumResponse, error) {
	response := &service_v1.SumResponse{
		Sum: req.First + req.Second,
	}

	return response, nil
}

func (s *server)ComputeAverage(stream service_v1.ComputeAverageService_ComputeAverageServer) error {
	var sum float64
	var numOfInputs int

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&service_v1.ComputeAverageResponse{
				Result: sum / float64(numOfInputs),
			})
			break
		}

		if err != nil {
			fmt.Printf("ComputeAverage: encountered error: %v", err)
			return err
		}

		fmt.Printf("Received %d\n", req.GetNumber())
		sum = sum + float64(req.GetNumber())
		numOfInputs = numOfInputs + 1
	}

	return nil
}

func main() {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatalf("Invalid server port")
	}

	address := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Printf("Starting server at %s\n", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service_v1.RegisterSumServiceServer(grpcServer, &server{})
	service_v1.RegisterComputeAverageServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
