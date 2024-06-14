package resolvers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/kelpi-io/kelpi-io/driver/storages"
	"github.com/redis/go-redis/v9"
)

func Lookup(qname string, qtype string, source string, config storages.Config, rdb *redis.Client) []RecordInfo {
	//ctx := context.Background()
	log.Println(source, qname)

	//con := rdb.Conn()

	//defer con.Close()

	//str := con.Get(ctx, "gslb-operator.io").Val()

	var ret []RecordInfo

	soa := getSOA(qname, config)

	ret = append(ret, soa...)

	return ret
}

func getSOA(qname string, config storages.Config) []RecordInfo {
	if qname == config.RootDomain {
		SOAstring := fmt.Sprintf(
			"%s %s %s %s %s %s %s",
			config.SoaMname,
			config.SoaRname,
			config.SoaSerial,
			config.SoaRefresh,
			config.SoaRetry,
			config.SoaExpire,
			config.SoaMinimum,
		)

		SoaTTLNum, _ := strconv.Atoi(config.SoaTTL)

		return []RecordInfo{
			{QType: "SOA", QName: qname, Content: SOAstring, TTL: SoaTTLNum},
		}
	} else {
		return []RecordInfo{}
	}
}
