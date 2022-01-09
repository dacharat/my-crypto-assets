package route

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/cmd/api/handler"
	"github.com/dacharat/my-crypto-assets/pkg/app"
	"github.com/gin-gonic/gin"
)

func NewRouter(app app.App) *gin.Engine {
	route := gin.Default()

	mid := newMiddleware(app.GetConfig().DevMode)
	h := handler.NewHandler(app.GetMyAssetsSvc(), app.GetLineSvc())

	route.GET("", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	route.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	route.GET("/test", mid.DevMode(), h.GetAccountBalanceHandler)
	route.POST("/linebot", h.LineCallbackHandler)
	route.GET("/push", mid.DevMode(), h.LinePushMessageHandler)

	return route
}
