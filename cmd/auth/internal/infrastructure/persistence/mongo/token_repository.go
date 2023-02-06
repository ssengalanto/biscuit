package mongo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ssengalanto/biscuit/cmd/auth/internal/domain/token"
	"github.com/ssengalanto/biscuit/pkg/errors"
	"github.com/ssengalanto/biscuit/pkg/interfaces"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TokenRepository - token repository struct.
type TokenRepository struct {
	log   interfaces.Logger
	token *mongo.Collection
}

const tokenCollection = "token"

// NewTokenRepository creates a new token repository instance.
func NewTokenRepository(log interfaces.Logger, db *mongo.Database) *TokenRepository {
	tkc := db.Collection(tokenCollection)
	return &TokenRepository{
		log:   log,
		token: tkc,
	}
}

// Create inserts a new Token record.
func (t *TokenRepository) Create(ctx context.Context, entity token.Entity) error {
	_, err := t.token.InsertOne(ctx, entity)
	if err != nil {
		t.log.Error("persisting token record failed", map[string]any{"payload": entity, "error": err})
		return fmt.Errorf("%w: presisting token failed", errors.ErrInternal)
	}

	return nil
}

// FindByID retrieve a Token record with the specified ID in the database.
func (t *TokenRepository) FindByID(ctx context.Context, id uuid.UUID) (token.Entity, error) {
	var tk Token
	filter := bson.D{bson.E{Key: "_id", Value: id.String()}}

	err := t.token.FindOne(ctx, filter).Decode(&tk)
	if err != nil {
		t.log.Error("retrieving token record failed", map[string]any{"id": id, "error": err})
		return tk.ToEntity(), fmt.Errorf("%w: token with id of `%s`", errors.ErrNotFound, id)
	}

	return tk.ToEntity(), nil
}

// DeleteByID deletes a Token record with the specified ID in the database.
func (t *TokenRepository) DeleteByID(ctx context.Context, id uuid.UUID) error {
	filter := bson.D{bson.E{Key: "_id", Value: id.String()}}

	res, err := t.token.DeleteOne(ctx, filter)
	if err != nil {
		t.log.Error("deleting token record failed", map[string]any{"id": id, "error": err})
		return fmt.Errorf("%w: token with id of `%s`", errors.ErrNotFound, id)
	}

	if res.DeletedCount <= 0 {
		return ErrDeleteRecordFailed
	}

	return nil
}
