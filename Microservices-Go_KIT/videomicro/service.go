package videomicro

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// // Declare the service interface and all the abstract methods inside it which you are going to
//  implement in the service layer.

// VideoService ...
type VideoService interface {
	CreateService(ctx context.Context, videomodel VideoModel) (string, error)
	GetService(ctx context.Context) (interface{}, error)
	UpdateService(ctx context.Context, id string, videomodel VideoModel) (string, error)
	DeleteService(ctx context.Context, id string) (string, error)
}

// // Also write a service struct and NewService which are useful when you are interacting
//  with the database through repository.
// Then implement all the service methods.

type videoServiceStruct struct {
	repository Repository
	logger     log.Logger
}

// NewService creates and returns a new User service instance

// NewService ...
func NewService(rep Repository, logger log.Logger) VideoService {
	return &videoServiceStruct{
		repository: rep,
		logger:     logger,
	}
}

func (vs videoServiceStruct) CreateService(ctx context.Context, videomodel VideoModel) (string, error) {
	logger := log.With(vs.logger, "method", "CreateService")
	videoDetails := VideoModel{
		Success: videomodel.Success,
		Data:    videomodel.Data,
	}
	if _, err := vs.repository.Create(ctx, videoDetails); err != nil {
		level.Error(logger).Log("err", err)
	}
	return "Done", nil
}

func (vs videoServiceStruct) GetService(ctx context.Context) (interface{}, error) {
	logger := log.With(vs.logger, "method", "GetService")
	var srv interface{}
	srv, err := vs.repository.Get(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return srv, nil
}

func (vs videoServiceStruct) UpdateService(ctx context.Context, id string, videomodel VideoModel) (string, error) {

	logger := log.With(vs.logger, "method", "UpdateService")
	_, err := vs.repository.Update(ctx, id, videomodel)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return "done", nil
}

func (vs videoServiceStruct) DeleteService(ctx context.Context, id string) (string, error) {
	logger := log.With(vs.logger, "method", "DeleteService")
	_, err := vs.repository.Delete(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return "Done", nil
}
