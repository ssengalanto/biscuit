package mongo

import (
	"context"
	"fmt"
	"net"

	"github.com/ssengalanto/biscuit/pkg/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewConnection initializes a new mongo database connection.
func NewConnection(user, pwd, host, port, dbname string, qp string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.ResourceTimeout)
	defer cancel()

	hp := net.JoinHostPort(host, port)
	dsn := fmt.Sprintf("mongodb://%s:%s@%s/?%s", user, pwd, hp, qp)
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	err = cl.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrConnectionFailed, err)
	}

	db := cl.Database(dbname)

	return db, nil
}
