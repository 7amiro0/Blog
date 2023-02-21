package main

import "os"

type Config struct {
	loggerLevel string
	mongoURI    string
	server      serverConfig
}

type serverConfig struct {
	host string
	port string
}

func (s *serverConfig) Set() {
	s.host = os.Getenv("HOST")
	s.port = os.Getenv("PORT")
}

func (c *Config) Set() {
	c.loggerLevel = os.Getenv("LEVEL")
	c.mongoURI = os.Getenv("URI")
	c.server.Set()
}

func NewConfig() Config {
	config := Config{}
	config.Set()

	return config
}
