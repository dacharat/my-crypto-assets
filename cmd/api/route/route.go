package route

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/cmd/api/handler"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	algoApi := algorand.NewAlgolandService()
	priceApi := coingecko.NewCoingeckoService()
	h := handler.NewHandler(algorandservice.NewService(algoApi, priceApi))
	route.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	route.GET("/test", h.GetAccountBalanceHandler)

	return route
}
