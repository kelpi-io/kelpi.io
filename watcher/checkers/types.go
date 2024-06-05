package checkers

// Params for all hosts in pool
type MonitorParam struct {
	Timeout  int `yaml:"timeout"`
	Interval int `yaml:"interval"`
	Retries  int `yaml:"retries"`
	Port     int `yaml:"port"`
	//	ReqData string
	//	RespRegex string
}

// Endpoint for monitoring
type Member struct {
	Name   string `yaml:"name"`
	Ip     string `yaml:"ip"`
	Weight int    `yaml:"weight"`
}

type WatcherConfig struct {
	GlobalName  string       `yaml:"globalName"`
	BalanceType string       `yaml:"balanceType"`
	Type        string       `yaml:"type"`
	Monitor     MonitorParam `yaml:"monitor"`
	Members     []Member     `yaml:"members"`
}
