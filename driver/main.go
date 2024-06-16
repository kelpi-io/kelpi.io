package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelpi-io/kelpi-io/driver/balancers"
	"github.com/kelpi-io/kelpi-io/driver/resolvers"
	"github.com/kelpi-io/kelpi-io/driver/storages"
)

func main() {

	// Read config
	storages.ConfigDriver.LoadConfig()
	config := &storages.ConfigDriver

	// Redis connect
	err := storages.Connect(config.RedisHost, config.RedisPassword, int(config.RedisDB))

	if err != nil {
		panic(err)
	}

	// Register Balancers

	balancers.Balancers = map[string]balancers.BalancerPrototype{
		"static": balancers.GetStatic}

	// Register HTTP handlers
	log.Println("Started driver for", config.RootDomain)

	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())

	server.GET("pdns/getAllDomains", resolvers.GetAllDomainsHandler())
	server.GET("pdns/lookup/:qname/:qtype", resolvers.LookupHandler())
	server.GET("pdns/getAllDomainMetadata/:qname", resolvers.GetAllDomainMetadataHandler())
	server.GET("pdns/initialize", resolvers.InitializeHandler())

	server.Run("0.0.0.0:8080")
}
