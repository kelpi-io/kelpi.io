package resolvers

import (
	"fmt"
	"strconv"

	"github.com/kelpi-io/kelpi-io/driver/balancers"
	"github.com/kelpi-io/kelpi-io/driver/storages"
)

func Lookup(qname string, qtype string, source string) []RecordInfo {

	// Get records
	var ret []RecordInfo
	recordSoa := getSOA(qname)
	recordA := getA(qname)

	// Prepare return array
	ret = append(ret, recordSoa...)
	ret = append(ret, recordA...)

	return ret
}

func getA(qname string) []RecordInfo {

	pool, err := storages.GetPool(qname)

	if err != nil {
		return []RecordInfo{}
	}

	currentWatcher := balancers.Balancers[pool.BalanceType]

	records := currentWatcher(pool)

	var ret []RecordInfo

	for _, val := range records {
		currentRecord := RecordInfo{
			QType:   "A",
			QName:   qname,
			Content: val,
			TTL:     1,
		}

		ret = append(ret, currentRecord)
	}

	return ret

	//log.Println(wc)
	// balancers.GetStatic(wc)

}

func getSOA(qname string) []RecordInfo {
	config := storages.ConfigDriver

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
