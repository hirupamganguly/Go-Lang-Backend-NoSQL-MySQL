package videomicro

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// VideoCollection ...
const VideoCollection = "videocols"

type repo struct {
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

func (repo *repo) Create(ctx context.Context, videomodel VideoModel) (interface{}, error) {
	collection := repo.client.Database("usermanagement").Collection(VideoCollection)
	result, err := collection.InsertOne(ctx, videomodel)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (repo *repo) Get(ctx context.Context) (interface{}, error) {
	var v []VideoModel
	collection := repo.client.Database("usermanagement").Collection(VideoCollection)
	cursor, _ := collection.Find(ctx, bson.M{})
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var vm VideoModel
		cursor.Decode(&vm)
		v = append(v, vm)
	}
	return v, nil
}

func (repo *repo) Delete(ctx context.Context, id string) (string, error) {
	collection := repo.client.Database("usermanagement").Collection(VideoCollection)
	nid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": nid}
	fmt.Println(filter)
	// d, err := collection.DeleteOne(ctx, VideoModel{ID: nid})
	d, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(d.DeletedCount)), nil
}

func (repo *repo) Update(ctx context.Context, id string, videomodel VideoModel) (string, error) {
	collection := repo.client.Database("usermanagement").Collection(VideoCollection)
	nid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": nid}
	upd := bson.D{
		{"$set", bson.D{
			{"success", videomodel.Success},
			{"data", videomodel.Data},
		}},
	}
	err := collection.FindOneAndUpdate(ctx, filter, upd).Decode(&videomodel)
	if err != nil {
		return "", err
	}
	return "Done", nil
}
