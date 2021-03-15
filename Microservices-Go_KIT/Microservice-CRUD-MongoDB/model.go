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

// ## business logic
//  The business logic in the services contain core business logic, which should not have any knowledge of endpoint or concrete transports like HTTP or gRPC, or encoding and decoding of request and response message types. This will encourage you follow a clean architecture for the Go kit based services
//  Each service method converts as an endpoint by using an adapter and exposed by using concrete transports.
//   ## endpoint
//   In Go kit, the primary messaging pattern is RPC. An endpoint represents a single RPC method. Each service method in a Go kit service converts to an endpoint to make RPC style communication between servers and clients. Each endpoint exposes the service method to outside world using Transport layer by using concrete transports like HTTP or gRPC. A single endpoint can be exposed by using multiple transports.
//   ## transport layer
//   The transport layer in Go kit is bound to concrete transports. Go kit supports various transports for serving services using HTTP, gRPC, NATS, AMQP and Thrift. Because Go kit services are just focused on implementing business logic and don’t have any knowledge about concrete transports, you can provide multiple transports for a same service.
