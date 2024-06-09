package checkers

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
	"time"
)

// Input param for monitor
type TCPMonitorParam struct {
	Timeout  int `json:"timeout"`
	Interval int `json:"interval"`
	Port     int `json:"port"`

	//	ReqData string
	//	RespRegex string
}

// Output param for monitor
type TcpHealthData struct {
	Health    bool   `json:"health"`
	LastCheck int64  `json:"lastCheck"`
	Latency   int64  `json:"latency"`
	IP        string `json:"ip"`
}

// Check healt host with TCP method
func TcpCheck(config WatcherConfig, memberName string) interface{} {

	// ============================
	// Parse params
	// ============================

	var monitorParam TCPMonitorParam
	errorParse := json.Unmarshal(config.Monitor, &monitorParam)

	if errorParse != nil {
		panic(errorParse)
	}

	member := config.Members[memberName]
	timeout := time.Duration(time.Second * time.Duration(monitorParam.Timeout))
	endpoint := net.JoinHostPort(member.Ip, strconv.Itoa(monitorParam.Port))

	// ============================
	// Main code
	// ============================

	startTime := time.Now()
	conn, err := net.DialTimeout("tcp", endpoint, timeout)
	endTime := time.Now()

	if err != nil {
		log.Printf("[%s] TCP error: %s", member.Ip, err)
	}

	defer conn.Close()

	health := TcpHealthData{
		Health:    err == nil,
		LastCheck: time.Now().Unix(),
		IP:        member.Ip,
		Latency:   endTime.Sub(startTime).Milliseconds(),
	}

	return health
}
