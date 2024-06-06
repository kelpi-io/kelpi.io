package checkers

import (
	"log"
	"net"
	"strconv"
	"time"
)

func TcpCheck(monitorParam MonitorParam, member Member) error {
	endpoint := net.JoinHostPort(
		member.Ip,
		strconv.Itoa(monitorParam.Port))

	timeout := time.Duration(time.Second * time.Duration(monitorParam.Timeout))

	conn, err := net.DialTimeout("tcp", endpoint, timeout)
	if err != nil {
		log.Printf("[%s] %s TCP error: %s", member.Name, member.Ip, err)
		return err
	}

	log.Printf("[%s] %s TCP OK", member.Name, member.Ip)

	defer conn.Close()

	return nil
}
