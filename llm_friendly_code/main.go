package main

import (
	"fmt"
	"log"
	"net"

	. "github.com/chaocai2001/llm_friendly_code/demo"
	"github.com/chaocai2001/llm_friendly_code/endpoint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Test: grpcui -plaintext localhost:8082

const Port = 8082

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}

// Assemble Processing Service
func createProcessingService() ProcessingService {
	processor := &ReverseProcessor{}
	tokenCreator := &MyTokenCreator{}
	storage := NewLocalStorage()
	return NewProcessingService(processor, tokenCreator, storage)
}

func startServer() {
	// Attach GRPC endpoint to ProcessingService
	endPointImpl := NewGRPC_Endpoint(createProcessingService())

	gRPCServer := grpc.NewServer()
	endpoint.RegisterProcessingServiceServer(gRPCServer, endPointImpl)
	reflection.Register(gRPCServer)
	log.Printf("Processing server is started ...")
	if err := gRPCServer.Serve(getNetListener(Port)); err != nil {
		panic(err)
	}
}

func main() {
	startServer()
}
