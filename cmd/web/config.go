package main

import "os"

type Config struct {
	loggerLevel string
	mongoURI    string
	server      serverConfig
	redis redisConfig
}

type serverConfig struct {
	host string
	port string
}

type redisConfig struct {
	addr string
	password string
}

func (r *redisConfig) Set() {
	r.addr = os.Getenv("REDIS_ADDR")
	r.password = os.Getenv("REDIS_PASSWORD")
}

func (s *serverConfig) Set() {
	s.host = os.Getenv("HOST")
	s.port = os.Getenv("PORT")
}

func (c *Config) Set() {
	c.loggerLevel = os.Getenv("LEVEL")
	c.mongoURI = os.Getenv("URI")
	c.redis.Set()
	c.server.Set()
}

func NewConfig() Config {
	config := Config{}
	config.Set()

	return config
}
