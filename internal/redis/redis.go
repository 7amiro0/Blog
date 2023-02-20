package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	queue *redis.Client
	ctx context.Context
}

func New(ctx context.Context, opt *redis.Options) *Redis {
	client := redis.NewClient(opt)

	return &Redis{
		queue: client,
		ctx: ctx,
	}
}

func (r *Redis) Add(key string, value interface{}) {
	r.queue.RPush(r.ctx, key, value)
}

func (r *Redis) Delete(key string, timeOut time.Duration) {
	r.queue.BLPop(r.ctx, timeOut, key)
}