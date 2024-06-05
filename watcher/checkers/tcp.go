package checkers

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type TCPMonitor struct {
	
}

func TcpCheck() error {
	endpoint := net.JoinHostPort("188.114.96.3", "80")
	timeout := time.Duration(time.Second * time.Duration(50))

	// if config.Verbose {
	// 	log.Println("runner: rawconnect: " + endpoint)
	// }

	// open the socket
	resp, err := net.DialTimeout("tcp", endpoint, timeout)
	if err != nil {
		log.Output(3, err.Error())
	}

	place := `
GET /todos/1 HTTP/1.1
User-Agent: PostmanRuntime/7.38.0
Accept: */*
Postman-Token: ed4a7ed8-aea8-4c06-bab6-0ece26627918
Host: jsonplaceholder.typicode.com
Accept-Encoding: gzip, deflate, br
Connection: keep-alive

	`

	log.Output(1, place)

	fmt.Fprint(resp, "GET /todos/1 HTTP/1.1\r\nHost: jsonplaceholder.typicode.com\r\n\r\n")

	reader := bufio.NewReader(resp)
	for reader.Size() != 0 {
		message, err := reader.ReadString('\n')
		fmt.Print("->: " + message)

		if err != nil {
			break
		}
	}

	resp.Close()
	//defer conn.Close()

	return nil
}
