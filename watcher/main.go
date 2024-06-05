package main

import (
	"github.com/vaishutin/gslb-operator/watcher/checkers"
)

func main() {

	checkers.TcpCheck()

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
