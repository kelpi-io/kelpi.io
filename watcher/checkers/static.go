package checkers

import (
	"encoding/json"
	"time"
)

// Input param for monitor
type StaticMonitorParam struct {
	Enabled bool `json:"enabled"`

	//	ReqData string
	//	RespRegex string
}

// Output param for monitor
type StaticData struct {
	Health    bool   `json:"health"`
	LastCheck int64  `json:"lastCheck"`
	Latency   int64  `json:"latency"`
	IP        string `json:"ip"`
}

// Check healt host with TCP method
func StaticCheck(config WatcherConfig, memberName string) interface{} {

	// ============================
	// Parse params
	// ============================

	var monitorParam StaticMonitorParam
	errorParse := json.Unmarshal(config.Monitor, &monitorParam)

	if errorParse != nil {
		panic(errorParse)
	}

	member := config.Members[memberName]

	// ============================
	// Main code
	// ============================

	health := StaticData{
		Health:    monitorParam.Enabled,
		LastCheck: time.Now().Unix(),
		IP:        member.Ip,
		Latency:   0,
	}

	return health
}
