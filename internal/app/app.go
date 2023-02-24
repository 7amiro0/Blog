package app

import (
	"blog/internal/storage"
)

type Logger interface {
	Fatal(msg ...any)
	Error(msg ...any)
	Debug(msg ...any)
	Warn(msg ...any)
	Info(msg ...any)
}

type Storage interface {
	Add(blog storage.Blog) error
	List() ([]storage.Blog, error)
	GetPost(title, author string) ([]storage.Blog, error)
}