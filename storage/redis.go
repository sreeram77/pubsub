package storage

import (
	"github.com/redis/go-redis/v9"
)

type cache struct {
	client *redis.Client
}

func New() Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &cache{
		client: rdb,
	}
}

func (c *cache) Get(key string) []any {
	return []any{}
}

func (c *cache) Set(key string, value any) {

}
