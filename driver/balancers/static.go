package balancers

import (
	"github.com/kelpi-io/kelpi-io/driver/storages"
)

func GetStatic(wc storages.WatcherConfig) []string {

	members := storages.GetMembers[storages.BaseMemberHealth](wc)

	var ret []string
	for _, data := range members {
		if !data.Health {
			continue
		}

		ret = append(ret, data.Ip)

	}

	return ret
}
