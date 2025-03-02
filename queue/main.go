package queue

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	queue *redis.Client
	once  sync.Once
)

func Queue() *redis.Client {
	once.Do(func() {
		queue = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		if err := queue.Ping(context.Background()).Err(); err != nil {
			log.Fatalln("Unable to connect to redis client", err.Error())
		}

		log.Println("Connected to redis")

	})
	return queue
}
