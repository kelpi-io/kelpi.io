package balancers

import (
	"github.com/kelpi-io/kelpi-io/driver/storages"
)

func GetStatic(wc storages.WatcherConfig) []string {

	var ret []string

	ret = append(ret, wc.Members["ip2"].Ip)

	return ret
}
