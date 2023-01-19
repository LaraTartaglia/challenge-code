package mongodb

import (
	"context"
	"fmt"

	"go.mod/internal/core/domain"
	"go.mod/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type itemRepository struct {
	conn *mongo.Client
}

func NewItemRepository(conn *mongo.Client) (ports.ItemRepository, error) {
	if conn == nil {
		return nil, fmt.Errorf("mongo connection cannot be nil")
	}

	return &itemRepository{conn: conn}, nil
}

func (i *itemRepository) SaveItem(ctx context.Context, itemToSave *domain.Item) error {

	collection := i.conn.Database(dbName).Collection("items")

	result, err := collection.InsertOne(ctx, itemToSave)
	if err != nil {
		return fmt.Errorf("error insert item: %w", err)
	}

	fmt.Printf("se inserto %v", result.InsertedID)

	return nil
}

func (i *itemRepository) GetItemById(ctx context.Context, itemToFind string) error {

	collection := i.conn.Database(dbName).Collection("items")

	err := collection.FindOne(ctx, bson.D{}).Decode(itemToFind)
	if err != nil {
		return err
	}

	return nil
}

func (i *itemRepository) GetItems(ctx context.Context, itemToFind string) error {

	collection := i.conn.Database(dbName).Collection("items")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return nil
}
