package balancers

import "github.com/kelpi-io/kelpi-io/driver/storages"

var Balancers map[string]BalancerPrototype

type BalancerPrototype func(wc storages.WatcherConfig) []string
