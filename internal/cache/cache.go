package cache

import (
	"blog/internal/storage"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	client *redis.Client
}

func New(addr, password string, db int) *Cache {
	client := redis.NewClient(&redis.Options{
		Addr:    addr,
		Password:           password,
		DB:                 db,
	})

	return &Cache{
		client: client,
	}
}

func (c *Cache) Set(key string, blog storage.Blogs, lifetime time.Duration) error {
	mBlogs, err := json.Marshal(blog)
	if err != nil {
		return err
	}

	return c.client.Set(key, mBlogs, lifetime).Err()
}

func (c *Cache) Get(key string) (storage.Blogs, error) {
	blogs := storage.Blogs{}
	values := c.client.Get(key)
	
	if values.Err() != nil {
		return blogs, values.Err()
	}

	bValue, err := values.Bytes()
	if err != nil {
		return blogs, err
	}

	err = json.Unmarshal(bValue, &blogs)
	if err != nil {
		return blogs, err
	}

	return blogs, nil
}