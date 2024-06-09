package checkers

import (
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"time"
)

// Input param for monitor
type HTTPMonitorParam struct {
	UseHttps   bool              `json:"https"`
	Host       string            `json:"host"`
	Path       string            `json:"path"`
	Headers    map[string]string `json:"headers"`
	Port       int               `json:"port"`
	ValidCodes []int             `json:"valid_codes"`

	Timeout  int `json:"timeout"`
	Interval int `json:"interval"`
}

// Output param for monitor
type HttpHealthData struct {
	Health    bool   `json:"health"`
	LastCheck int64  `json:"lastCheck"`
	IP        string `json:"ip"`
	Status    string `json:"status"`
	Latency   int64  `json:"latency"`
}

// Check healt host with TCP method
func HttpCheck(config WatcherConfig, memberName string) interface{} {

	// ============================
	// Parse params
	// ============================

	var monitorParam HTTPMonitorParam
	errorParse := json.Unmarshal(config.Monitor, &monitorParam)

	if errorParse != nil {
		panic(errorParse)
	}

	member := config.Members[memberName]
	timeout := time.Duration(time.Second * time.Duration(monitorParam.Timeout))
	url := urlJoin(member, monitorParam)

	// ============================
	// Main code
	// ============================

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Host = monitorParam.Host

	req.Header.Set("User-Agent", "gslb-operator/1.0.0")

	for k, v := range monitorParam.Headers {
		req.Header.Set(k, v)
	}

	startTime := time.Now()
	res, errRes := client.Do(req)
	endTime := time.Now()

	data := HttpHealthData{
		Health:    false,
		LastCheck: time.Now().Unix(),
		IP:        member.Ip,
		Latency:   endTime.Sub(startTime).Milliseconds(),
	}

	if errRes == nil {
		if slices.Contains(monitorParam.ValidCodes, res.StatusCode) {
			data.Health = true
		}

		data.Status = res.Status
	} else {
		data.Status = errRes.Error()
	}

	return data
}

func urlJoin(member Member, monParam HTTPMonitorParam) string {
	endpoint := net.JoinHostPort(member.Ip, strconv.Itoa(monParam.Port))

	url := url.URL{
		Host: endpoint,
		Path: monParam.Path,
	}

	if monParam.UseHttps {
		url.Scheme = "https"
	} else {
		url.Scheme = "http"
	}

	return url.String()
}
