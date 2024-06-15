package storages

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
)

// Init connect
func GetClientRDB(addr string, password string, db int) error {
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
		return cmd.Err()
	}
	log.Println("[REDIS]", cmd)

	RDB = client

	return nil
}

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

func GetPool(conn *redis.Conn, qname string) WatcherConfig {
	ctx := context.Background()
	ret := (*RDB).Get(ctx, qname)
	if ret.Err() != nil {
		panic(ret.Err())
	}

	var wconf WatcherConfig

	if err := json.Unmarshal([]byte(ret.Val()), &wconf); err != nil {
		panic(err)
	}

	return wconf

}

func GetMembers(conn *redis.Conn, wc WatcherConfig) []Member {
	ctx := context.Background()

	var keys []string

	for member := range wc.Members {
		keys = append(keys, wc.GlobalName+"/"+member)
	}

	ret := conn.MGet(ctx, keys...)

	if ret.Err() != nil {
		panic(ret.Err())
	}

	return []Member{}

}
