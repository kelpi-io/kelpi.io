package checkers

import (
	"log"
	"net"
	"strconv"
	"time"
)

type TcpHealthData struct {
	Health    bool   `json:"health"`
	LastCheck int64  `json:"lastCheck"`
	IP        string `json:"ip"`
}

func TcpCheck(config WatcherConfig, member Member) interface{} {
	endpoint := net.JoinHostPort(
		member.Ip,
		strconv.Itoa(config.Monitor.Port))

	timeout := time.Duration(time.Second * time.Duration(config.Monitor.Timeout))

	conn, err := net.DialTimeout("tcp", endpoint, timeout)

	if err != nil {
		log.Printf("[%s] TCP error: %s", member.Ip, err)
	}

	defer conn.Close()

	health := TcpHealthData{
		Health:    err == nil,
		LastCheck: time.Now().Unix(),
		IP:        member.Ip,
	}

	return health
}
