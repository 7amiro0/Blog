package app

import (
	"blog/internal/storage"
	"time"
)

type Logger interface {
	Fatal(msg ...any)
	Error(msg ...any)
	Debug(msg ...any)
	Warn(msg ...any)
	Info(msg ...any)
}

type Cache interface {
	Set(key string, blog storage.Blogs, lifetime time.Duration) error
	Get(key string) (storage.Blogs, error)
}

type Storage interface {
	Add(blog storage.Blog) error
	List() (storage.Blogs, error)
	GetPost(title, author string) (storage.Blogs, error)
}