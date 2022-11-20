package mongo

import (
	"context"
	"fmt"
	"restApi/internal/domain/user"
	"restApi/pkg/logging"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/genproto/googleapis/cloud/resourcemanager/v2"
)

type database struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (db *database) Create(ctx context.Context, user user.User) (string, error) {
	db.logger.Debug("create user")
	result, err := db.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	db.logger.Debug("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	db.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex")
}

func (db *database) FindOne(ctx context.Context, id string) (u user.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to objectid. hex: %s", id)
	}
	filter := bson.M{"_id": oid}

	result := db.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return u, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user (id:%s) from DB due to error: %v", id, err)
	}

	return u, nil
}

func (db *database) Update(ctx context.Context, user user.User) error {

}

func (db *database) Delete(ctx context.Context, id string) error {

}

func NewStorage(md *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &database{
		collection: md.Collection(collection),
		logger:     logger,
	}
}
