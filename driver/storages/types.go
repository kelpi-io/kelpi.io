package storages

import "encoding/json"

type Config struct {
	RootDomain    string
	RedisHost     string
	RedisPassword string
	RedisDB       int
	SoaMname      string
	SoaRname      string
	SoaSerial     string
	SoaRefresh    string
	SoaRetry      string
	SoaExpire     string
	SoaMinimum    string
	SoaTTL        string
}

type BaseMemberHealth struct {
	Ip        string `json:"ip"`
	Health    bool   `json:"health"`
	LastCheck int    `json:"lastCheck"`
	Latency   int    `json:"latency"`
	Status    string `json:"status"`
}

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
