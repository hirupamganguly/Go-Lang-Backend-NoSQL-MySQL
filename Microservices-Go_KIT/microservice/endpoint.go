package mymicroservice

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	StatusEndpoint  endpoint.Endpoint
	GetDataEndpoint endpoint.Endpoint
}

// Endpoints have a simple definition. They take in a context and a request, and return a response and an error.

//MakeStatusEndpoint ...
func MakeStatusEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(statusReq)
		s, err := srv.Status(ctx)
		if err != nil {
			return statusRes{s}, err
		}
		return statusRes{s}, nil
	}
}

// MakeGetDataEndpoint ...
func MakeGetDataEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getDataReq)
		g, err := srv.GetData(ctx)
		if err != nil {
			return getDataRes{g, err.Error()}, nil
		}
		return getDataRes{g, ""}, nil
	}
}
