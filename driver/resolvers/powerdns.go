package resolvers

import (
	"fmt"
	"strconv"

	"github.com/kelpi-io/kelpi-io/driver/balancers"
	"github.com/kelpi-io/kelpi-io/driver/storages"
	"github.com/redis/go-redis/v9"
)

func Lookup(qname string, qtype string, source string, config storages.Config, rdb *redis.Client) []RecordInfo {
	// Prepare connect
	con := rdb.Conn()
	defer con.Close()

	// Get records
	var ret []RecordInfo
	recordSoa := getSOA(qname, config)
	recordA := getA(qname, con)

	// Prepare return array
	ret = append(ret, recordSoa...)
	ret = append(ret, recordA...)

	return ret
}

func getA(qname string, conn *redis.Conn) []RecordInfo {
	wc := storages.GetPool(conn, qname)

	balancers.GetStatic(conn, wc)
	return []RecordInfo{}

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
