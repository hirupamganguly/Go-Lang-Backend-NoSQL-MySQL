package mymicroservice

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints ...
type Endpoints struct {
	StatusEndpoint       endpoint.Endpoint
	GetDataEndpoint      endpoint.Endpoint
	ValidateDataEndpoint endpoint.Endpoint
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

// MakeValidateDataEndpoint ...
func MakeValidateDataEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateDataReq)
		v, err := srv.ValidateData(ctx, req.Date)
		if err != nil {
			return validateDataRes{v, err.Error()}, err
		}
		return validateDataRes{v, ""}, nil
	}
}

// // Get ...
// func (e Endpoints) Get(ctx context.Context) (string, error) {
// 	req := getDataReq{}
// 	resp, err := e.GetDataEndpoint(ctx, req)
// 	if err != nil {
// 		return "", err
// 	}
// 	getDataRes := resp.(getDataRes)
// 	if getDataRes.Err != "" {
// 		return "", errors.New(getDataRes.Err)
// 	}
// 	return getDataRes.Date, nil
// }

// // Status ...
// func (e Endpoints) Status(ctx context.Context) (string, error) {
// 	req := statusReq{}
// 	resp, err := e.StatusEndpoint(ctx, req)
// 	if err != nil {
// 		return "", err
// 	}
// 	statusRes := resp.(statusRes)
// 	return statusRes.Status, nil
// }

// // Validate ...
// func (e Endpoints) Validate(ctx context.Context, date string) (bool, error) {

// 	req := validateDataReq{Date: date}
// 	resp, err := e.ValidateDataEndpoint(ctx, req)
// 	if err != nil {
// 		return false, err
// 	}
// 	validateDataRes := resp.(validateDataRes)
// 	if validateDataRes.Err != "" {
// 		return false, errors.New(validateDataRes.Err)
// 	}
// 	return validateDataRes.Valid, nil

// }
