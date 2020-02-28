package api

import (
	"github.com/gin-gonic/gin"
	"github.com/theojulienne/go-wireless"
)

func notImplemented(c *gin.Context) {
	c.AbortWithStatus(501)
}

func listInterfaces(c *gin.Context) {
	c.JSON(200, wireless.Interfaces())
}

func listAccesPoints(c *gin.Context) {
	iface := c.Param("iface")
	wc, err := wireless.NewClient(iface)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, json(err))
	}

	aps, err := wc.Scan()

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(500, json(err))
	}

	c.JSON(200, aps)
}

func json(err) map[string]string {
	return map[string]string{"error": err.Error()}
}
