package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/kelpi-io/kelpi-io/watcher/checkers"
	"github.com/kelpi-io/kelpi-io/watcher/storage"
	"github.com/redis/go-redis/v9"
)

func main() {

	// ============================
	// Read params
	// ============================

	poolConfig := flag.String("config", "./pool-config.json", "Path to config file with filename")
	redisHost := flag.String("rhost", "localhost:6379", "Redis host")
	redisPassword := flag.String("rpass", "qwerty", "Redis password")
	redisDB := flag.Int("rdb", 0, "Redis DB number")

	flag.Parse()
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
	// Registration checkers
	// ============================

	checkersMap := map[string]checkers.CheckerPrototype{
		"tcp":    checkers.TcpCheck,
		"http":   checkers.HttpCheck,
		"static": checkers.StaticCheck,
	}

	currentChecker := checkersMap[configs.MonitorType]

	// ============================
	// Go run gorutines
	// ============================

	waitGroup := &sync.WaitGroup{}

	for memberName := range configs.Members {
		waitGroup.Add(1)
		go worker(configs, memberName, waitGroup, client.Conn(), currentChecker)
	}

	waitGroup.Wait()

}

func worker(config checkers.WatcherConfig, memberName string, wg *sync.WaitGroup, conn *redis.Conn, f checkers.CheckerPrototype) {
	defer wg.Done()

	log.Printf("[checker] Started for %s", memberName)

	for {

		result := f(config, memberName)

		_ = storage.WriteStat(conn, config, memberName, result)

		time.Sleep(time.Second * time.Duration(config.Interval))

	}

}
