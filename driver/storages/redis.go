package storages

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

// Init connect
func GetClient(addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: 1000,
	})

	ctx := context.Background()

	cmd := client.Ping(ctx)
	if cmd.Err() != nil {
		log.Println("Redis connect error")
		return nil, cmd.Err()
	}
	log.Println("[REDIS]", cmd)

	return client, nil
}
