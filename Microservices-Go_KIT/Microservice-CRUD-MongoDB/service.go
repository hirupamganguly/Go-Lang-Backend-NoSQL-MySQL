package crudmicro

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Declare the service interface and all the abstract methods inside it which you are going to
//  implement in the service layer.
// Service describes the User service.

// UserService ...
type UserService interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUserByID(ctx context.Context, id string) (interface{}, error)
	GetAllUsers(ctx context.Context) (interface{}, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	UpdateUser(ctx context.Context, id string, user User) error
}

// Also write a service struct and NewService which are useful when you are interacting
//  with the database through repository.
// Then implement all the service methods.
// NewService creates and returns a new User service instance

type userserviceStruct struct {
	repository Repository
	logger     log.Logger // logger for logging middleware
}

// NewService ...
func NewService(rep Repository, logger log.Logger) UserService {
	return &userserviceStruct{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an user
func (s userserviceStruct) CreateUser(ctx context.Context, user User) (string, error) {
	// syntax: func With(logger Logger, keyvals ...interface{}) Logger
	logger := log.With(s.logger, "method", "Create")
	userDetails := User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
	if err := s.repository.CreateUser(ctx, userDetails); err != nil {
		level.Error(logger).Log("err", err) // Error returns a logger that includes a Key/ErrorValue pair.
	}
	return userDetails.ID.String(), nil
}

func (s userserviceStruct) GetUserByID(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetUserById")
	var data interface{}
	data, err := s.repository.GetUserByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)

		return "", err
	}
	return data, nil
}

func (s userserviceStruct) GetAllUsers(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllUsers")
	var email interface{}
	email, err := s.repository.GetAllUsers(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}
	return email, nil
}

func (s userserviceStruct) DeleteUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteUser")
	msg, err := s.repository.DeleteUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)

		return "", err
	}
	return msg, nil
}

func (s userserviceStruct) UpdateUser(ctx context.Context, id string, user User) error {
	logger := log.With(s.logger, "method", "ChangeDetails")
	// var email string
	err := s.repository.UpdateUser(ctx, id, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}
