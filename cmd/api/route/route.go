package route

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/cmd/api/handler"
	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/binance"
	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/external/line"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/binanceservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	client, err := linebot.New(config.Cfg.Line.ChannelSecret, config.Cfg.Line.ChannelAccessToken)
	if err != nil {
		panic(err)
	}

	hc := httpclient.NewClient()

	algoApi := algorand.NewAlgolandService(hc)
	priceApi := coingecko.NewCoingeckoService(hc)
	bitkubApi := bitkub.NewBitkubService(hc)
	binancApi := binance.NewBinanceService(hc)
	lineApi := line.NewLineService(client)

	algoSvc := algorandservice.NewService(algoApi, priceApi)
	bitkubSvc := bitkubservice.NewService(bitkubApi)
	binanceSvc := binanceservice.NewService(binancApi)
	myAssetsSvc := myassetsservice.NewService(algoSvc, bitkubSvc, binanceSvc)
	lineSvc := lineservice.NewService(lineApi)

	h := handler.NewHandler(myAssetsSvc, lineSvc, client.ParseRequest)

	route.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	route.GET("/test", DevMode(), h.GetAccountBalanceHandler)
	route.POST("/linebot", h.LineCallbackHandler)
	route.GET("/push", DevMode(), h.LinePushMessageHandler)

	return route
}
