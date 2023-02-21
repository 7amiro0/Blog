package storage

import (
	"blog/internal/logger"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	ctx context.Context
	client *mongo.Client
}

var logg *logger.Logger

func New(uri string, ctx context.Context, log *logger.Logger) (*Storage, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	logg = log
	
	return &Storage{
		ctx: ctx,
		client: client,
	}, nil
}

func (s *Storage) Connect() error {
	return s.client.Connect(s.ctx)
}

func (s *Storage) Disconnect() error {
	return s.client.Disconnect(s.ctx)
}

func (s *Storage) Add(blog Blog) error {
	collection := s.client.Database("blogs").Collection("blogs")
	_, err := collection.InsertOne(s.ctx, 
		bson.M{"title": blog.Title, "author": blog.Author, "body": blog.Body, "views": 0},
	)
	
	return err
}

func getPosts(values *mongo.Cursor) []Blog {
	result := make([]Blog, 0, 1)

	for values.Next(context.Background()) {
		var blog Blog
		err := values.Decode(&blog)
		if err != nil {
			logg.Info("Error to decode blog: ", err)
		}

		result = append(result, blog)
	}

	return result
}

func (s *Storage) List() ([]Blog, error) {
	collection := s.client.Database("blogs").Collection("blogs")
	values, err := collection.Find(s.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer values.Close(s.ctx)

	return getPosts(values), nil
}

func (s *Storage) IncreaseViews(title, author string) error {
	collection := s.client.Database("blogs").Collection("blogs")
	_, err := collection.UpdateOne(s.ctx, bson.M{"title": title, "author": author}, bson.M{"$inc": bson.M{"views": 1}})
	return err
}



func (s *Storage) GetPost(title, author string) ([]Blog, error) {
	collection := s.client.Database("blogs").Collection("blogs")
	values, err := collection.Find(s.ctx, bson.M{"title": title, "author": author})
	if err != nil {
		return nil, err
	}
	defer values.Close(s.ctx)

	return getPosts(values), nil
}