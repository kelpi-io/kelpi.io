package resolvers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelpi-io/kelpi-io/driver/storages"
)

func GetAllDomainsHandler() func(c *gin.Context) {

	res := Response{
		Result: []DomainInfoResult{
			{Zone: storages.ConfigDriver.RootDomain},
		},
	}

	return func(c *gin.Context) {

		c.JSON(200, res)
	}

}

func LookupHandler() func(c *gin.Context) {

	return func(c *gin.Context) {

		qname := c.Param("qname")
		qtype := c.Param("qtype")
		fwdForHeaders := c.Request.Header[http.CanonicalHeaderKey("X-Remotebackend-Remote")]
		fwdFor := "127.0.0.1"

		if len(fwdForHeaders) > 0 {
			fwdFor = fwdForHeaders[0]
		}
		//log.Println(fwdFor, qname)

		ri := Lookup(qname, qtype, fwdFor)

		if len(ri) > 0 {
			res := Response{
				Result: ri,
			}

			c.JSON(200, res)
		} else {
			res := Response{
				Result: false,
			}

			c.JSON(200, res)
		}

		//
		//fwdFor := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-for")]
		//fwdPort := c.Request.Header[http.CanonicalHeaderKey("x-forwarded-port")]

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
