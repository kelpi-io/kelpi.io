package checkers

import (
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"net/url"
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
	Retries  int `json:"retries"`
}

// Output param for monitor
type HttpHealthData struct {
	Health     bool   `json:"health"`
	LastCheck  int64  `json:"lastCheck"`
	IP         string `json:"ip"`
	StatusCode string `json:"statusCode"`
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

	req.Header.Add("User-Agent", "gslb-operator/1.0.0")

	for k, v := range monitorParam.Headers {
		req.Header.Add(k, v)
	}

	res, errRes := client.Do(req)

	health := HttpHealthData{
		Health:    errRes == nil,
		LastCheck: time.Now().Unix(),
		IP:        member.Ip,
	}

	if res != nil {
		health.StatusCode = res.Status
	}

	return health
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
