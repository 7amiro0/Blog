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

type Queue interface {
	Add(key string, value interface{})
	Delete(key string, timeOut time.Duration)
}

type StorageBlogs interface {
	Add(blog storage.Blog) error
}

type StorageAnalitics interface {
	Add() error
	Update() error
	Get() error
}