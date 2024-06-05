package main

import (
	"time"

	"github.com/vaishutin/gslb-operator/watcher/checkers"
)

func main() {

	mon := checkers.MonitorParam{
		Timeout:  1,
		Port:     80,
		Interval: 2,
	}

	member := checkers.Member{
		Name:   "test",
		Ip:     "188.114.96.3",
		Weight: 1,
	}

	for {
		_ = checkers.TcpCheck(mon, member)

		time.Sleep(time.Second * time.Duration(mon.Interval))
	}

	// client := http.Client{}

	// ctx, cancel := context.WithTimeout(context.Background(), 10000)
	// defer cancel()

	// req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://10.10.10.91", nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if _, err := client.Do(req); err != nil {
	// 	log.Fatal(err)
	// }
}
