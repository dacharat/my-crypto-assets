package app

import (
	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/binance"
	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/external/elrond"
	"github.com/dacharat/my-crypto-assets/pkg/external/line"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/binanceservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubchainservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/elrondservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type App struct {
	myAssetsSvc    myassetsservice.IMyAssetsService
	lienSvc        lineservice.ILineService
	planetwatchSvc platnetwatchservice.IPlanetwatchService
	cfg            *config.Config
}

func New(cfg *config.Config) App {
	client, err := linebot.New(cfg.Line.ChannelSecret, cfg.Line.ChannelAccessToken)
	if err != nil {
		panic(err)
	}

	hc := httpclient.NewClient()
	conn, _ := ethclient.Dial(cfg.ChainRpc.Bitkub)

	algoApi := algorand.NewAlgolandService(hc, &cfg.AlgorandClient)
	priceApi := coingecko.NewCoingeckoService(hc, &cfg.Coingecko)
	bitkubApi := bitkub.NewBitkubService(hc, &cfg.Bitkub)
	binancApi := binance.NewBinanceService(hc, &cfg.Binance)
	elrondApi := elrond.NewService(hc, &cfg.Elrond)
	lineApi := line.NewLineService(client, &cfg.Line)

	assetsServices := []shared.IAssetsService{
		algorandservice.NewService(algoApi, &cfg.AlgorandClient),
		bitkubservice.NewService(bitkubApi),
		binanceservice.NewService(binancApi),
		bitkubchainservice.NewService(conn),
		elrondservice.NewService(elrondApi),
	}

	planetwatchSvc := platnetwatchservice.NewService(algoApi, &cfg.AlgorandClient, cfg.User.AlgoAddress)
	myAssetsSvc := myassetsservice.NewService(assetsServices, priceApi, &cfg.User)
	lineSvc := lineservice.NewService(lineApi, &cfg.User, cfg.Line.UserID)

	return App{
		myAssetsSvc:    myAssetsSvc,
		lienSvc:        lineSvc,
		planetwatchSvc: planetwatchSvc,
		cfg:            cfg,
	}
}

func (a App) GetMyAssetsSvc() myassetsservice.IMyAssetsService {
	return a.myAssetsSvc
}

func (a App) GetLineSvc() lineservice.ILineService {
	return a.lienSvc
}

func (a App) GetPlanetwatchSvc() platnetwatchservice.IPlanetwatchService {
	return a.planetwatchSvc
}

func (a App) GetConfig() *config.Config {
	return a.cfg
}
