package crudmicro

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User ...
type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             ` json:"email,omitempty" bson:"email,omitempty"`
	Password string             ` json:"password,omitempty"  bson:"password,omitempty"`
	Phone    string             ` json:"phone,omitempty"  bson:"phone,omitempty"`
}

// Repository ...
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByID(ctx context.Context, id string) (interface{}, error)
	GetAllUsers(ctx context.Context) (interface{}, error)
	DeleteUser(ctx context.Context, id string) (string, error)
	UpdateUser(ctx context.Context, id string, user User) error
}
