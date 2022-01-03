package route

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/cmd/api/handler"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/binance"
	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/binanceservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubservice"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	algoApi := algorand.NewAlgolandService()
	priceApi := coingecko.NewCoingeckoService()
	bitkubApi := bitkub.NewBitkubService()
	binancApi := binance.NewBinanceService()

	algoSvc := algorandservice.NewService(algoApi, priceApi)
	bitkubSvc := bitkubservice.NewService(bitkubApi)
	binanceSvc := binanceservice.NewService(binancApi)

	h := handler.NewHandler(algoSvc, bitkubSvc, binanceSvc)

	route.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	route.GET("/test", h.GetAccountBalanceHandler)

	return route
}
