package app

import (
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
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type App struct {
	myAssetsSvc myassetsservice.IMyAssetsService
	lienSvc     lineservice.ILineService
}

func New() App {
	client, err := linebot.New(config.Cfg.Line.ChannelSecret, config.Cfg.Line.ChannelAccessToken)
	if err != nil {
		panic(err)
	}

	hc := httpclient.NewClient()
	_, _ = ethclient.Dial("https://rpc.bitkubchain.io")

	algoApi := algorand.NewAlgolandService(hc)
	priceApi := coingecko.NewCoingeckoService(hc)
	bitkubApi := bitkub.NewBitkubService(hc)
	binancApi := binance.NewBinanceService(hc)
	lineApi := line.NewLineService(client)

	assetsServices := []shared.IAssetsService{
		algorandservice.NewService(algoApi, priceApi),
		bitkubservice.NewService(bitkubApi),
		binanceservice.NewService(binancApi),
	}

	myAssetsSvc := myassetsservice.NewService(assetsServices, priceApi)
	lineSvc := lineservice.NewService(lineApi)

	return App{
		myAssetsSvc: myAssetsSvc,
		lienSvc:     lineSvc,
	}
}

func (a App) GetMyAssetsSvc() myassetsservice.IMyAssetsService {
	return a.myAssetsSvc
}

func (a App) GetLineSvc() lineservice.ILineService {
	return a.lienSvc
}
