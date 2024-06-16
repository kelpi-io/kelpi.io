package storages

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client = RedisRepository{}

type RedisRepository struct {
	Connection *redis.Client
	Ctx        context.Context
}

func Connect(addr string, password string, db int) error {
	Client.Connection = redis.NewClient(&redis.Options{
		Addr:            addr,
		Password:        password,
		DB:              db,
		PoolSize:        1000,
		ConnMaxLifetime: (1 * time.Minute),
	})

	Client.Ctx = context.Background()

	conn := Client.Connection.Conn()
	defer conn.Close()

	cmd := conn.Ping(Client.Ctx)
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

func GetMembers[T interface{}](wc WatcherConfig) []T {

	var keys []string

	for member := range wc.Members {
		fullKey := fmt.Sprintf("%s/%s/health", wc.GlobalName, member)
		keys = append(keys, fullKey)
	}

	conn := Client.Connection.Conn()
	defer conn.Close()

	redisCmd := conn.MGet(Client.Ctx, keys...)

	if redisCmd.Err() == redis.Nil {
		return []T{}
	} else if redisCmd.Err() != nil {
		panic(redisCmd.Err())
	}

	var ret []T

	for _, data := range redisCmd.Val() {
		var memberHealth T
		if err := json.Unmarshal([]byte(data.(string)), &memberHealth); err != nil {
			log.Println(err)
			continue
		}

		ret = append(ret, memberHealth)
	}

	return ret

}
