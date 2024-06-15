package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelpi-io/kelpi-io/driver/resolvers"
	"github.com/kelpi-io/kelpi-io/driver/storages"
)

func main() {
	config := storages.LoadConfig()

	storages.GetClientRDB(config.RedisHost, config.RedisPassword, int(config.RedisDB))

	rdb, err := storages.GetClient(config.RedisHost, config.RedisPassword, int(config.RedisDB))

	if err != nil {
		panic(err)
	}

	log.Println("Started driver for", config.RootDomain)

	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())

	server.GET("pdns/getAllDomains", resolvers.GetAllDomainsHandler(&config))
	server.GET("pdns/lookup/:qname/:qtype", resolvers.LookupHandler(config, rdb))
	server.GET("pdns/getAllDomainMetadata/:qname", resolvers.GetAllDomainMetadataHandler())
	server.GET("pdns/initialize", resolvers.InitializeHandler())

	server.Run("0.0.0.0:8080")
}
