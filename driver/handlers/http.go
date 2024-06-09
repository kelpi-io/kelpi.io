package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

func getAllDomainsHandler() func(c *gin.Context) {

	res := Response{
		Result: []DomainInfoResult{
			{Zone: "www.example.com."},
		},
	}

	return func(c *gin.Context) {

		c.JSON(200, res)
	}

}

func lookupHandler() func(c *gin.Context) {

	res := Response{
		Result: []RecordInfo{
			{QType: "SOA", QName: "www.example.com", Content: "gslb.example.com. hostmaster.polaris.example.com. 1 3600 600 86400 1", TTL: 1},
			{QType: "A", QName: "www.example.com", Content: "203.0.113.2", TTL: 1},
		},
	}

	return func(c *gin.Context) {
		qname := c.Param("qname")
		qtype := c.Param("qtype")

		log.Println(qname)
		log.Println(qtype)
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdPort := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-port")]

		c.JSON(200, res)
	}

}

func getAllDomainMetadataHandler() func(c *gin.Context) {

	res := Response{
		Result: map[string]interface{}{
			"PRESIGNED": []string{"0"},
		},
	}

	return func(c *gin.Context) {
		qname := c.Param("qname")
		qtype := c.Param("qtype")

		log.Println(qname)
		log.Println(qtype)
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdPort := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-port")]

		c.JSON(200, res)
	}

}

func initializeHandler() func(c *gin.Context) {

	res := Response{
		Result: true,
	}

	return func(c *gin.Context) {
		qname := c.Param("qname")
		qtype := c.Param("qtype")

		log.Println(qname)
		log.Println(qtype)
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdPort := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-port")]

		c.JSON(200, res)
	}

}

func StartHttpHandler() {
	server := gin.Default()
	server.GET("pdns/getAllDomains", getAllDomainsHandler())
	server.GET("pdns/lookup/:qname/:qtype", lookupHandler())
	server.GET("pdns/getAllDomainMetadata/:qname", getAllDomainMetadataHandler())
	server.GET("pdns/initialize", initializeHandler())

	server.Run("0.0.0.0:8080")
}
