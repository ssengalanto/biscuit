package mongo

import (
	"context"

	"github.com/ssengalanto/biscuit/pkg/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewConnection initializes a new mongo database connection.
func NewConnection(dsn string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ResourceTimeout)
	defer cancel()

	db, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, ErrConnectionFailed
	}

	err = db.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, ErrConnectionFailed
	}

	return db, nil
}
