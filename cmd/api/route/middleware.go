package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func newMiddleware(devMode bool) middleware {
	return middleware{
		devMode: devMode,
	}
}

type middleware struct {
	devMode bool
}

func (m middleware) DevMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !m.devMode {
			c.AbortWithStatus(http.StatusTeapot)
			return
		}
		c.Next()
	}
}
