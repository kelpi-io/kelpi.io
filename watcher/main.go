package main

import (
	"flag"
	"log"
	"os"
	"sync"
	"time"

	"github.com/vaishutin/gslb-operator/watcher/checkers"
	"gopkg.in/yaml.v3"
)

func main() {

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
		go func(mon checkers.MonitorParam, member checkers.Member, wg *sync.WaitGroup) {
			defer wg.Done()
			log.Printf("[checker] Started for %s", member.Ip)
			for {

				_ = checkers.TcpCheck(
					mon,
					member)

				time.Sleep(time.Second * time.Duration(mon.Interval))
			}

		}(configs.Monitor, configs.Members[memberIndex], waitGroup)
	}

	waitGroup.Wait()

	log.Println("Done")

}
