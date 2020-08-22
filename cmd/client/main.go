package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"google.golang.org/grpc"
	service_v1 "grpc-exercises/service/v1"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	sumServiceLabel                = "Sum"
	primeDecompositionServiceLabel = "Prime Number Decomposition"
	computeAverageServiceLabel     = "Compute Average"
	quitLabel                      = "Quit"
)

func main() {
	serverUrl := os.Getenv("SERVER_URL")
	if serverUrl == "" {
		log.Fatalf("SERVER_URL is required")
	}

	for {
		prompt := promptui.Select{
			Label: "Select service to connect to",
			Items: []string{
				sumServiceLabel,
				primeDecompositionServiceLabel,
				computeAverageServiceLabel,
				quitLabel,
			},
		}

		_, selectedService, err := prompt.Run()

		if err != nil {
			log.Fatalf("Failed to get user input for service to connect to: %v\n", err)
		}

		if selectedService == quitLabel {
			break
		} else {
			startClient(serverUrl, selectedService)
		}
	}
}

func startClient(serverUrl string, selectedService string) {
	fmt.Println("==================================")
	clientConnection, err := grpc.Dial(serverUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s with error: %v", serverUrl, err)
	}

	defer clientConnection.Close()

	fmt.Printf("[%s] Connected to server at: %s\n", selectedService, serverUrl)

	switch selectedService {
	case sumServiceLabel:
		startSumClient(clientConnection)
		break
	case primeDecompositionServiceLabel:
		startPrimeDecompositionClient(clientConnection)
		break
	case computeAverageServiceLabel:
		startComputeAverageClient(clientConnection)
		break
	}
	fmt.Println("==================================")
}

func startPrimeDecompositionClient(connection *grpc.ClientConn) {
	client := service_v1.NewPrimeNumberDecompositionServiceClient(connection)

	input, err := getIntInput("Enter a number to decompose")
	if err == io.EOF {
		log.Fatalf("No user input, exiting")
	}

	stream, err := client.Decompose(context.Background(), &service_v1.PrimeNumberDecompositionRequest{
		Input: input,
	})

	if err != nil {
		log.Fatalf("Error while calling server: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Server finished sending")
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving result from server: %v", err)
		}

		fmt.Printf("Result: %v\n", res.GetResult())
	}
}

func startComputeAverageClient(connection *grpc.ClientConn) {
	client := service_v1.NewComputeAverageServiceClient(connection)
	stream, err := client.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling server: %v", err)
	}

	for {
		input, err := getIntInput("Enter an integer")
		if err == io.EOF {
			res, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("Failed to receive response from server: %v", err)
			} else {
				fmt.Printf("=== Response from server: %v\n", res.GetResult())
			}

			break
		}

		stream.Send(&service_v1.ComputeAverageRequest{Number: input})
		fmt.Printf("=== Sent %v to server\n", input)
	}
}

func startSumClient(connection *grpc.ClientConn) {
	client := service_v1.NewSumServiceClient(connection)

	first, err := getIntInput("Enter first number")
	if err == io.EOF {
		log.Fatalf("No user input, exiting")
	}

	second, err := getIntInput("Enter second number")
	if err == io.EOF {
		log.Fatalf("No user input, exiting")
	}

	request := &service_v1.SumRequest{
		First:  first,
		Second: second,
	}

	response, err := client.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling server: %v", err)
	}

	fmt.Printf("Result: %d + %d = %d\n", first, second, response.Sum)
}

func getIntInput(label string) (int32, error) {
	validate := func(input string) error {
		if input == "" {
			return nil
		}

		_, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			return errors.New("invalid integer provided")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	selected, err := prompt.Run()

	if selected == "" {
		return 0, io.EOF
	}

	if err != nil {
		return 0, errors.New(fmt.Sprintf("Failed to get user input: %v", err))
	}

	parsedResult, _ := strconv.ParseInt(selected, 10, 32)
	return int32(parsedResult), nil
}
