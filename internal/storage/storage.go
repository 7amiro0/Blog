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
		bson.M{"title": blog.Title, "author": blog.Author, "body": blog.Body},
	)
	
	return err
}

func getPosts(values *mongo.Cursor) Blogs {
	result := Blogs{
		Blogs: make([]Blog, 0, 1),
	}

	for values.Next(context.Background()) {
		var blog Blog
		err := values.Decode(&blog)
		if err != nil {
			logg.Info("Error to decode blog: ", err)
		}

		result.Blogs = append(result.Blogs, blog)
	}

	return result
}

func (s *Storage) List() (Blogs, error) {
	collection := s.client.Database("blogs").Collection("blogs")
	values, err := collection.Find(s.ctx, bson.M{})
	if err != nil {
		return Blogs{nil}, err
	}
	defer values.Close(s.ctx)

	return getPosts(values), nil
}

func (s *Storage) GetPost(title, author string) (Blogs, error) {
	collection := s.client.Database("blogs").Collection("blogs")
	values, err := collection.Find(s.ctx, bson.M{"title": title, "author": author})
	if err != nil {
		return Blogs{nil}, err
	}
	defer values.Close(s.ctx)

	return getPosts(values), nil
}