package resolvers

import (
	"github.com/gin-gonic/gin"
	"github.com/kelpi-io/kelpi-io/driver/storages"
	"github.com/redis/go-redis/v9"
)

func GetAllDomainsHandler(config *storages.Config) func(c *gin.Context) {

	res := Response{
		Result: []DomainInfoResult{
			{Zone: config.RootDomain},
		},
	}

	return func(c *gin.Context) {

		c.JSON(200, res)
	}

}

func LookupHandler(config storages.Config, rdb *redis.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		qname := c.Param("qname")
		qtype := c.Param("qtype")
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]

		ri := Lookup(qname, qtype, "fwdFor[0]", config, rdb)

		res := Response{
			Result: ri,
		}

		//
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdPort := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-port")]

		c.JSON(200, res)
	}

}

func GetAllDomainMetadataHandler() func(c *gin.Context) {

	res := Response{
		Result: map[string]interface{}{
			"PRESIGNED": []string{"0"},
		},
	}

	return func(c *gin.Context) {
		c.JSON(200, res)
	}

}

func InitializeHandler() func(c *gin.Context) {

	res := Response{
		Result: true,
	}

	return func(c *gin.Context) {
		c.JSON(200, res)
	}

}
