package videomicro

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// VideoDataModel ...
type VideoDataModel struct {
	VideoID        string `json:"video_id" bson:"video_id"`
	VideoTitle     string `json:"video_title" bson:"video_title"`
	VideoThumbnail string `json:"video_thumbnail" bson:"video_thumbnail"`
	VideoType      string `json:"video_type" bson:"video_type"`
	Like           int    `json:"like" bson:"like"`
}

// CollectionModel ...
type CollectionModel struct {
	CollectionID   string           `json:"collection_id" bson:"collection_id"`
	CollectionName string           `json:"collection_name" bson:"collection_name"`
	NosOfVideo     int              `json:"nos_of_video" bson:"nos_of_video"`
	Video          []VideoDataModel `json:"video" bson:"video"`
}

// Data ...
type Data struct {
	Data []CollectionModel `json:"collection" bson:"collection"`
}

// VideoModel ...
type VideoModel struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Success bool               `json:"success" bson:"success"`
	Data    Data               `json:"data" bson:"data"`
}

// Repository ...
type Repository interface {
	Create(ctx context.Context, videomodel VideoModel) (interface{}, error)
	Get(ctx context.Context) (interface{}, error)
	Update(ctx context.Context, id string, videomodel VideoModel) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}
