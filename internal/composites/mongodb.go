package composites

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"restful_go_project/pkg/client/mongodb"
)

type MongoDBComposite struct {
	db *mongo.Database
}

func NewMongoDBComposite(ctx context.Context, Host, Port, Username, Password, Database, AuthDB string) (*MongoDBComposite, error) {
	client, err := mongodb.NewClient(ctx, Host, Port, Username, Password, Database, AuthDB)
	if err != nil {
		return nil, err
	}

	return &MongoDBComposite{
		db: client,
	}, nil
}
