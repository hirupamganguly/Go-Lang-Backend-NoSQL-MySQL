package videomicro

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	CreateEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}

// MakeCreateEndpoint ...
func MakeCreateEndpoint(s VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		msg, err := s.CreateService(ctx, req.v)
		return CreateResponse{Msg: msg, Err: err}, nil
	}
}

// MakeGetEndpoint ...
func MakeGetEndpoint(s VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		data, err := s.GetService(ctx)
		return GetResponse{Data: data, Err: err}, nil
	}
}

// MakeUpdateEndpoint ...
func MakeUpdateEndpoint(s VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		_, err := s.UpdateService(ctx, req.ID, req.v)
		return UpdateResponse{Err: err}, nil

	}
}

// MakeDeleteEndpoint ...
func MakeDeleteEndpoint(s VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteRequest)
		id := req.ID
		_, err := s.DeleteService(ctx, id)
		return UpdateResponse{Err: err}, nil
	}
}
