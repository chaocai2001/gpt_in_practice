package demo

import (
	"context"

	"github.com/chaocai2001/llm_friendly_code/endpoint"
)

type GRPC_Endpoint struct {
	endpoint.UnimplementedProcessingServiceServer
	processingService ProcessingService
}

func NewGRPC_Endpoint(processingService ProcessingService) *GRPC_Endpoint {
	return &GRPC_Endpoint{
		processingService: processingService,
	}
}

func (ep *GRPC_Endpoint) Process(ctx context.Context, req *endpoint.ProcessingRequest) (*endpoint.ProcessingReply, error) {
	token, err := ep.processingService.Process(req.GetData())
	if err != nil {
		return nil, err
	}
	return &endpoint.ProcessingReply{Token: token}, nil
}

func (ep *GRPC_Endpoint) Retrieve(ctx context.Context, req *endpoint.RetrievingRequest) (*endpoint.RetrievingReply, error) {
	data, err := ep.processingService.Retrieve(req.GetToken())
	if err != nil {
		return nil, err
	}

	return &endpoint.RetrievingReply{Data: data}, nil
}
