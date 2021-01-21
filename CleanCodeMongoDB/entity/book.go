package entity

// Book ...
type Book struct {
	Name   string `json:"name" bson:"name"`
	Author string `json:"author" bson:"author"`
	ID     int64  `json:"id" bson:"id"`
}
