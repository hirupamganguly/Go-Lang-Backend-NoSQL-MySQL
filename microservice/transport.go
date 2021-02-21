package mymicroservice

import (
	"context"
	"encoding/json"
	"net/http"
)

// The transport.go is responsible for converting a transport layer request into an endpoint call.
type statusReq struct{}

type statusRes struct {
	Status string `json:"status"`
}

type getDataReq struct{}

type getDataRes struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

type validateDataReq struct {
	Date string `json:"date"`
}
type validateDataRes struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

func decodeGetDataReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getDataReq
	return req, nil
}
func decodeValidateDataReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req validateDataReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func decodeStatusReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req statusReq
	return req, nil
}

func encodeRes(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
