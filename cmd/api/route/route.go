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
	h := handler.NewHandler(app.GetMyAssetsSvc(), app.GetLineSvc(), app.GetPlanetwatchSvc())

	route.GET("", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})
	route.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	route.POST("/linebot", h.LineCallbackHandler)

	// For Development or use ngrok
	route.GET("/test", mid.DevMode(), h.GetAccountBalanceHandler)
	route.GET("/push", mid.DevMode(), h.LinePushMessageHandler)
	route.GET("/push-platform", mid.DevMode(), h.LinePushMessageByPlatformHandler)

	return route
}
