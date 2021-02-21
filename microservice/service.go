package mymicroservice

import (
	"context"
	"fmt"
	"time"
)

// Service ...
type Service interface {
	Status(ctx context.Context) (string, error)
	GetData(ctx context.Context) (string, error)
	ValidateData(ctx context.Context, item string) (bool, error)
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

func (dateService) ValidateData(ctx context.Context, item string) (bool, error) {
	tm, e := time.Parse("02/01/21", item)
	fmt.Println(tm)
	if e != nil {
		return false, e
	}
	return true, nil
}
