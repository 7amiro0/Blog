package main

import (
	"blog/internal/logger"
	"blog/internal/server"
	"blog/internal/storage"
	"context"
	"net"
	"os/signal"
	"syscall"
)

func main() {
	config := NewConfig()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	log := logger.New(config.loggerLevel)

	store, err := storage.New(config.mongoURI, ctx, log)
	if err != nil {
		log.Fatal(err)
	}

	if err = store.Connect(); err != nil {
		log.Fatal(err)
	}
	defer store.Disconnect()

	serv := server.New(net.JoinHostPort(config.server.host, config.server.port), log, store)

	log.Info("Servre has been created")

	go func() {
		if err = serv.Start(); err != nil {
			log.Error("Cant start server: ", err)
		} 
	}()
	defer serv.Stop(ctx)
	
	log.Info("Servre has been ran")

	<-ctx.Done()

	log.Info("Servre has been stoped")
}