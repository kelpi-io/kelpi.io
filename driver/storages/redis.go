package storages

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client = RedisRepository{}

type RedisRepository struct {
	Connection *redis.Client
	Ctx        context.Context
}

func Connect(addr string, password string, db int) error {
	Client.Connection = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		PoolSize: 1000,
	})

	Client.Ctx = context.Background()

	cmd := Client.Connection.Get(context.TODO(), "Test")
	if cmd.Err() != nil {
		log.Println("Redis connect error")
		return cmd.Err()
	}
	log.Println("[REDIS-connect]", cmd)

	return nil
}

func GetPool(qname string) (WatcherConfig, error) {

	conn := Client.Connection.Conn()
	defer conn.Close()

	ret := conn.Get(Client.Ctx, qname)

	if ret.Err() == redis.Nil {
		return WatcherConfig{}, errors.New("pool not found")
	} else if ret.Err() != nil {
		panic(ret.Err())
	}

	var wconf WatcherConfig

	if err := json.Unmarshal([]byte(ret.Val()), &wconf); err != nil {
		panic(err)
	}

	return wconf, nil

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
