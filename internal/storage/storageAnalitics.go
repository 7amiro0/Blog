package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageAnalitics struct {
	ctx context.Context
	client *mongo.Client
}

func NewStorageAnalitics(uri string, ctx context.Context) (*StorageAnalitics, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	
	return &StorageAnalitics{
		ctx: ctx,
		client: client,
	}, nil
}

func (s *StorageAnalitics) Connect() error {
	return s.client.Connect(s.ctx)
}

func (s *StorageAnalitics) Disconnect() error {
	return s.client.Disconnect(s.ctx)
}