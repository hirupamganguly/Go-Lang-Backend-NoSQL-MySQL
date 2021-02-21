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

func TestValidateData(t *testing.T) {
	srv, ctx := setup()
	st, _ := srv.ValidateData(ctx, "19/02/2021")
	if !st {
		t.Errorf("date should be valid")
	}
	st, _ = srv.ValidateData(ctx, "02/19/2021")
	if st {
		t.Errorf("date should be invalid")
	}
	st, _ = srv.ValidateData(ctx, "32/42/2021")
	if st {
		t.Errorf("date should be inavalid")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
