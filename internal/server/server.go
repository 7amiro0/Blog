package server

import (
	"blog/internal/app"
	"context"

	"net/http"
)

type Server struct {
	router *http.ServeMux
	server *http.Server
	cache app.Cache
	log app.Logger
	storage app.Storage
}

func New(addr string, log app.Logger, storage app.Storage, cache app.Cache) *Server {
	router := http.NewServeMux()

	server := &Server{
		router: router,
		log:    log,
		storage: storage,
		cache: cache,
	}

	server.setRouter()

	server.server = &http.Server{
		Addr: addr,
		Handler: server.router,
	}

	return server
}

func (s *Server) setRouter() {
	s.router.HandleFunc("/home", s.index)
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}