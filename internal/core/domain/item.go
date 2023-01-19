package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Code        string             `bson:"code"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       int                `bson:"price"`
	Stock       int                `bson:"stock"`
	Photos      []string           `bson:"photos"`
	CreatedAt   string             `bson:"created_at"`
	UpdatedAt   string             `bson:"updated_at"`
	Status      string             `bson:"status"`
}
