package mymicroservice

import (
	"context"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()
	st, _ := srv.Status(ctx)
	if st != "ok" {
		t.Errorf("Expected service to be ok")
	}
}

func TestGetData(t *testing.T) {
	srv, ctx := setup()
	st, _ := srv.GetData(ctx)
	time := time.Now()
	today := time.Format("02/01/2001")
	if today != st {
		t.Errorf("Expected dates to be equal")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
