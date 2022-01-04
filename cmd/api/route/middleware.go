package route

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/gin-gonic/gin"
)

func DevMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.Cfg.DevMode {
			c.AbortWithStatus(http.StatusTeapot)
			return
		}
		c.Next()
	}
}
