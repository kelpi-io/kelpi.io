package storage

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/vaishutin/gslb-operator/watcher/checkers"
)

// Init connect
func GetClient(addr string, password string, db int, poolName string) (*redis.Client, error) {
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

// Write Global Name struct to redis
func InitPool(conn *redis.Conn, conf checkers.WatcherConfig) error {
	ctx := context.Background()

	data, _ := json.Marshal(conf)

	cmd := conn.Set(ctx, conf.GlobalName, data, 0)

	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return cmd.Err()
	}

	return nil
}

// Write member status
func WriteStat(conn *redis.Conn, config checkers.WatcherConfig, memberName string, healthData interface{}) error {
	ctx := context.Background()

	keyVal := config.GlobalName + "/" + memberName + "/health"
	value, _ := json.Marshal(healthData)

	cmd := conn.Set(ctx, keyVal, value, 0)

	if cmd.Err() != nil {
		panic(cmd.Err())
	}

	return cmd.Err()

}
