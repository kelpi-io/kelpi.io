package resolvers

import (
	"context"

	"github.com/kelpi-io/kelpi-io/driver/storages"
	"github.com/redis/go-redis/v9"
)

func Lookup(qname string, qtype string, source string, config storages.Config, rdb *redis.Client) []RecordInfo {
	ctx := context.Background()

	con := rdb.Conn()

	defer con.Close()

	str := con.Get(ctx, "gslb-operator.io").Val()

	return []RecordInfo{
		{QType: str},
	}
}

// res := Response{
// 	Result: []RecordInfo{
// 		{QType: "SOA", QName: "www.example.com", Content: "gslb.example.com. hostmaster.polaris.example.com. 1 3600 600 86400 1", TTL: 1},
// 		{QType: "A", QName: "www.example.com", Content: "203.0.113.2", TTL: 1},
// 	},
// }
