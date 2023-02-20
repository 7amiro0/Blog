package main

import "os"

type Config struct {
	loggerLevel string
	mongoURI    string
	server      serverConfig
	redis       redisConfig
}

type serverConfig struct {
	addres string
}

type redisConfig struct {
	addres string
}

func (r *redisConfig) Set() {
	r.addres = os.Getenv("REDISADDR")
}

func (s *serverConfig) Set() {
	s.addres = os.Getenv("SERVERADDR")
}

func (c *Config) Set() {
	c.loggerLevel = os.Getenv("LEVEL")
	c.mongoURI = os.Getenv("URI")
	c.server.Set()
	c.redis.Set()
}

func NewConfig() Config {
	config := Config{}
	config.Set()

	return config
}
