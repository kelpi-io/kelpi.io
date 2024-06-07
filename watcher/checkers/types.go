package checkers

// Params for all hosts in pool
type MonitorParam struct {
	Timeout  int `json:"timeout"`
	Interval int `json:"interval"`
	Retries  int `json:"retries"`
	Port     int `json:"port"`
	//	ReqData string
	//	RespRegex string
}

// Endpoint for monitoring
type Member struct {
	Ip     string `json:"ip"`
	Weight int    `json:"weight"`
}

type WatcherConfig struct {
	GlobalName  string            `json:"globalName"`
	BalanceType string            `json:"balanceType"`
	Type        string            `json:"type"`
	Monitor     MonitorParam      `json:"monitor"`
	Members     map[string]Member `json:"members"`
}
