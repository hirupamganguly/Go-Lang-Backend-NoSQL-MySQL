package crudmicro

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserCollection ...
const UserCollection = "usercols"

type repo struct {
	// Client is a handle representing a pool of connections to a MongoDB deployment. It is safe for concurrent use by multiple goroutines.
	// The Client type opens and closes connections automatically and maintains a pool of idle connections. For connection pool configuration options, see documentation for the ClientOptions type in the mongo/options package.

	client *mongo.Client
	logger log.Logger
	err    error
}

// NewRepo ...
func NewRepo(client *mongo.Client, logger log.Logger) (Repository, error) {
	return &repo{
		client: client,
		logger: log.With(logger, "repo", "mongodb"),
	}, nil
}

// --------THE Repository contract is present inside model.go--------

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	collection := repo.client.Database("usermanagement").Collection(UserCollection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Error occured inside CREATE USER in REPO")
		return err
	}
	fmt.Println("User Created:", user.Email)
	return nil
}
func (repo *repo) GetUserByID(ctx context.Context, id string) (interface{}, error) {
	collection := repo.client.Database("usermanagement").Collection(UserCollection)
	data := User{}
	nid, _ := primitive.ObjectIDFromHex(id)
	collection.FindOne(ctx, User{ID: nid}).Decode(&data)
	return data, nil
}

func (repo *repo) GetAllUsers(ctx context.Context) (interface{}, error) {
	var users []User
	collection := repo.client.Database("usermanagement").Collection(UserCollection)
	cursor, _ := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) { // The cursor.Next() method is used to return the next document in a cursor.
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}
func (repo *repo) DeleteUser(ctx context.Context, id string) (string, error) {
	collection := repo.client.Database("usermanagement").Collection(UserCollection)
	nid, _ := primitive.ObjectIDFromHex(id)
	collection.DeleteOne(ctx, User{ID: nid})
	return "deleted successfully", nil
}

func (repo *repo) UpdateUser(ctx context.Context, id string, user User) error {
	collection := repo.client.Database("usermanagement").Collection(UserCollection)
	nid, _ := primitive.ObjectIDFromHex(id)
	upd := bson.D{
		{"$set", bson.D{
			{"email", user.Email},
			{"password", user.Password},
			{"phone", user.Phone},
		}},
	}

	collection.FindOneAndUpdate(ctx, User{ID: nid}, upd).Decode(&user)

	return nil
}
