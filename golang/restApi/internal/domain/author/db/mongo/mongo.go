package mongo

import (
	"context"
	"errors"
	"fmt"
	"restApi/internal/custome_errors"
	"restApi/internal/domain/author"
	"restApi/pkg/logging"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type database struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (db *database) Create(ctx context.Context, a author.Author) (string, error) {
	db.logger.Debug("create author")
	result, err := db.collection.InsertOne(ctx, a)
	if err != nil {
		return "", fmt.Errorf("failed to create author due to error: %v", err)
	}

	db.logger.Debug("convert InsertedID to ObjectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	db.logger.Trace(a)
	return "", fmt.Errorf("failed to convert objectid to hex")
}

func (db *database) FindOne(ctx context.Context, id string) (a author.Author, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return a, fmt.Errorf("failed to convert hex to objectid. hex: %s", id)
	}
	filter := bson.M{"_id": oid}

	result := db.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return a, custome_errors.ErrNotFound
		}
		return a, fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
	}
	if err = result.Decode(&a); err != nil {
		return a, fmt.Errorf("failed to decode user (id:%s) from DB due to error: %v", id, err)
	}

	return a, nil
}

func (db *database) Update(ctx context.Context, a author.Author) error {
	objectID, err := primitive.ObjectIDFromHex(a.ID)
	if err != nil {
		return fmt.Errorf("failed to convert author ID to ObjectID. ID=%s", a.ID)
	}

	filter := bson.M{"_id": objectID}

	authorBytes, err := bson.Marshal(a)
	if err != nil {
		return fmt.Errorf("failed to marhsal user. error: %v", err)
	}

	var updateAuthorObj bson.M
	err = bson.Unmarshal(authorBytes, &updateAuthorObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal author bytes. error: %v", err)
	}

	delete(updateAuthorObj, "_id")

	update := bson.M{
		"$set": updateAuthorObj,
	}

	result, err := db.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to execute update author query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return custome_errors.ErrNotFound
	}

	db.logger.Tracef("Matched %d documents and Modified %d documents", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (db *database) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", id)
	}

	filter := bson.M{"_id": objectID}

	result, err := db.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %v", err)
	}
	if result.DeletedCount == 0 {
		return custome_errors.ErrNotFound
	}
	db.logger.Tracef("Delete %d documents", result.DeletedCount)

	return nil
}

func NewStorage(md *mongo.Database, collection string, logger *logging.Logger) author.Repository {
	return &database{
		collection: md.Collection(collection),
		logger:     logger,
	}
}
