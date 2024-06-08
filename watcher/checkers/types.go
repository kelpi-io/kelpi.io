package checkers

import "encoding/json"

type CheckerPrototype func(WatcherConfig, string) interface{}

// Endpoint for monitoring
type Member struct {
	Ip     string `json:"ip"`
	Weight int    `json:"weight"`
}

type WatcherConfig struct {
	GlobalName  string            `json:"globalName"`
	BalanceType string            `json:"balanceType"`
	Monitor     json.RawMessage   `json:"monitor"`
	Members     map[string]Member `json:"members"`
	MonitorType string            `json:"monitorType"`
	Interval    int               `json:"interval"`
}
