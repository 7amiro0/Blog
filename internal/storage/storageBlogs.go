package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageBlogs struct {
	ctx context.Context
	client *mongo.Client
}

func NewStorageBlogs(uri string, ctx context.Context) (*StorageBlogs, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	
	return &StorageBlogs{
		ctx: ctx,
		client: client,
	}, nil
}

func (s *StorageBlogs) Connect() error {
	return s.client.Connect(s.ctx)
}

func (s *StorageBlogs) Disconnect() error {
	return s.client.Disconnect(s.ctx)
}

func (s *StorageBlogs) Add(blog Blog) error {
	collection := s.client.Database("blogs").Collection("blogs")
	_, err := collection.InsertOne(s.ctx, 
		bson.M{"title": blog.Title, "author": blog.Author, "body": blog.Body},
	)
	
	return err
}