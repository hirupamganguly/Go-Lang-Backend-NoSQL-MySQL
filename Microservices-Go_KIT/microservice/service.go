package mymicroservice

import (
	"context"
	"time"
)

// Service ...
type Service interface {
	Status(ctx context.Context) (string, error)
	GetData(ctx context.Context) (string, error)
}
type dateService struct{}

// NewService makes a new Service.
func NewService() Service {
	return dateService{}
}

func (dateService) Status(ctx context.Context) (string, error) {
	return "ok", nil
}
func (dateService) GetData(ctx context.Context) (string, error) {
	now := time.Now()
	return now.Format("02/01/21"), nil
}
