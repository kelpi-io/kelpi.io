package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vaishutin/gslb-operator/watcher/checkers"
	"github.com/vaishutin/gslb-operator/watcher/storage"
)

func main() {

	// ============================
	// Read params
	// ============================

	poolConfig := flag.String("config", "./pool-config.json", "Path to config file with filename")
	redisHost := flag.String("rhost", "localhost:6379", "Redis host")
	redisPassword := flag.String("rpass", "qwerty", "Redis password")
	redisDB := flag.Int("rdb", 0, "Redis DB number")

	// ============================
	// Read and parse config from Json file
	// ============================

	configs, err := GetConfig(*poolConfig)

	if err != nil {
		panic(err)
	}

	// ============================
	// Connect to Redis
	// ============================

	client, _ := storage.GetClient(*redisHost, *redisPassword, *redisDB, configs.GlobalName)
	errRedis := storage.InitPool(client.Conn(), configs)
	defer client.Close()

	if errRedis != nil {
		panic(errRedis)
	}

	// ============================
	// Go run gorutines
	// ============================

	waitGroup := &sync.WaitGroup{}

	for memberName := range configs.Members {
		waitGroup.Add(1)
		go worker(configs, memberName, waitGroup, client.Conn())
	}

	waitGroup.Wait()

}

func worker(config checkers.WatcherConfig, memberName string, wg *sync.WaitGroup, conn *redis.Conn) {
	defer wg.Done()

	member := config.Members[memberName]
	log.Printf("[checker] Started for %s", member.Ip)

	for {

		err := checkers.TcpCheck(config.Monitor, member)

		_ = storage.WriteStat(conn, config, memberName, err == nil)

		time.Sleep(time.Second * time.Duration(config.Monitor.Interval))

	}

}
