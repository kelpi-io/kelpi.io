package main

import (
	"context"
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vaishutin/gslb-operator/watcher/checkers"
	"gopkg.in/yaml.v3"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "qwerty",
		DB:       0,
		PoolSize: 1000,
	})

	ctx := context.Background()

	cmd := client.Ping(ctx)
	if cmd.Err() != nil {
		log.Println("Redis connect error")
		panic(cmd.Err())
	}

	log.Println(cmd)

	cn := client.Conn()

	cn.Set(ctx, "first", "val", 5)

	// ============================
	// Read params
	// ============================

	poolConfig := flag.String("config", "./pool-config.yml", "Path to config file with filename")

	// ============================
	// Read and parse config from Yaml file
	// ============================
	yamlFile, err := os.ReadFile(*poolConfig)

	if err != nil {
		panic(err)
	}

	var configs checkers.WatcherConfig

	err = yaml.Unmarshal(yamlFile, &configs)

	if err != nil {
		panic(err)
	}

	// ============================
	// Go run gorutines
	// ============================

	waitGroup := &sync.WaitGroup{}

	for memberIndex := range configs.Members {

		log.Printf("[checker] Starting... for %s", configs.Members[memberIndex].Ip)

		waitGroup.Add(1)

		go func(mon checkers.MonitorParam, member checkers.Member, wg *sync.WaitGroup, conn *redis.Conn, ctx context.Context) {
			defer wg.Done()
			defer conn.Close()
			log.Printf("[checker] Started for %s", member.Ip)

			conn2 := client.Conn()
			for {

				err = checkers.TcpCheck(
					mon,
					member)

				memberEnd := member
				memberEnd.Health = err == nil
				memberEnd.LastCheck = time.Now().Unix()

				yml, _ := yaml.Marshal(memberEnd)

				status := conn2.Set(ctx, memberEnd.Ip, yml, 0)
				if status.Err() != nil {
					log.Println("Error write to Redis")
				}

				time.Sleep(time.Second * time.Duration(mon.Interval))

			}

		}(configs.Monitor, configs.Members[memberIndex], waitGroup, client.Conn(), context.Background())
	}

	waitGroup.Wait()

	log.Println("Done")

}
