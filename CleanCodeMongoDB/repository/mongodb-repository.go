package repository

import (
	"context"
	"time"

	"../entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewMongoRepository ...
func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (BookRepositoryContract, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (mr *mongoRepository) Find(name string) (*entity.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	book := &entity.Book{}
	collection := mr.client.Database(mr.database).Collection("books")
	filter := bson.M{"name": name}
	err := collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return book, nil
}
func (mr *mongoRepository) Store(book *entity.Book) error {
	ctx, cancel := context.WithTimeout(context.Background(), mr.timeout)
	defer cancel()
	collection := mr.client.Database(mr.database).Collection("books")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"name":   book.Name,
			"author": book.Author,
			"id":     book.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
