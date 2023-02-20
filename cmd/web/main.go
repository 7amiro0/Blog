package main

import (
	"blog/internal/logger"
	"blog/internal/redis"
	"blog/internal/server"
	"blog/internal/storage"
	"context"

	anr "github.com/redis/go-redis/v9"
)

func main() {
	config := NewConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log := logger.New(config.loggerLevel)
	
	queue := redis.New(ctx, &anr.Options{
		Addr:       config.redis.addres,
	})

	store, err := storage.NewStorageBlogs(config.mongoURI, ctx)
	if err != nil {
		log.Fatal(err)
	}

	if err = store.Connect(); err != nil {
		log.Fatal(err)
	}
	defer store.Disconnect()

	serv := server.New(config.server.addres, log, queue, store)

	log.Info("Servre has been created")

	go func() {
		if err := serv.Start(); err != nil {
			log.Error("Cant start server: ", err)
		} 
	}()
	defer serv.Stop(ctx)
	
	log.Info("Servre has been ran")

	<-ctx.Done()

	
	log.Info("Servre has been stoped")
}