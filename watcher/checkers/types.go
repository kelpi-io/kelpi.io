package checkers

// Params for all hosts in pool
type MonitorParam struct {
	Timeout  int
	Interval int
	Retries  int
	Port     int
	//	ReqData string
	//	RespRegex string
}

// Endpoint for monitoring
type Member struct {
	Name   string
	Ip     string
	Weight int
}
