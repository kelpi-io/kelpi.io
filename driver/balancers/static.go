package balancers

import (
	"github.com/kelpi-io/kelpi-io/driver/storages"
	"github.com/redis/go-redis/v9"
)

func GetStatic(conn *redis.Conn, wc storages.WatcherConfig) []string {

	storages.GetMembers(conn, wc)

	return []string{}
}
