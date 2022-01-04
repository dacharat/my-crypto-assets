package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Cfg config

type config struct {
	AlgorandClient Algorand
	Coingecko      Coingecko
	Bitkub         Bitkub
	Binance        Binance
	Line           Line
	User           User
	DevMode        bool `envconfig:"DEV_MODE" default:"false"`
}

type Algorand struct {
	Host         string `envconfig:"ALGORAND_HOST" required:"true"`
	GetAssetPath string `envconfig:"ALGORAND_GET_ASSET_PATH" default:"/v2/assets/%d"`

	AlgodHost      string `envconfig:"ALGORAND_ALGOD_HOST" required:"true"`
	ApiKey         string `envconfig:"ALGORAND_ALGOD_APILEY"`
	GetAccountPath string `envconfig:"ALGORAND_GET_ACCOUNT_PATH" default:"/v2/accounts/%s"`

	DefaultDecimal int  `envconfig:"ALGORAND_DEFAULT_DECIMAL" default:"6"`
	UseFreeApi     bool `envconfig:"ALGORAND_USE_FREE_API" default:"false"`
}

type Coingecko struct {
	Host           string `envconfig:"COINGECKGO_HOST" required:"true"`
	GetSimplePrice string `envconfig:"COINGECKGO_GET_SIMPLE_PRICE" default:"/api/v3/simple/price"`
}

type Bitkub struct {
	Host       string `envconfig:"BITKUB_HOST" required:"true"`
	ApiKey     string `envconfig:"BITKUB_APIKEY" required:"true"`
	ApiSecret  string `envconfig:"BITKUB_APISECRET" required:"true"`
	GetWallet  string `envconfig:"BITKUB_GET_WALLET" default:"/api/market/wallet"`
	GetTricker string `envconfig:"BITKUB_GET_TRICKER" default:"/api/market/ticker"`
}

type Binance struct {
	Host       string `envconfig:"BINANCE_HOST" required:"true"`
	ApiKey     string `envconfig:"BINANCE_APIKEY" required:"true"`
	ApiSecret  string `envconfig:"BINANCE_APISECRET" required:"true"`
	GetAccount string `envconfig:"BINANCE_GET_ACCOUNT" default:"/api/v3/account"`
	GetSaving  string `envconfig:"BINANCE_GET_SAVING" default:"/sapi/v1/lending/union/account"`
	GetTricker string `envconfig:"BINANCE_GET_TRICKER" default:"/api/v3/ticker/price"`
}

type Line struct {
	UserID             string `envconfig:"LINE_USER_ID" required:"true"`
	ChannelSecret      string `envconfig:"LINE_CHANNEL_SECRET" required:"true"`
	ChannelAccessToken string `envconfig:"LINE_CHANNEL_ACCESS_TOKEN" required:"true"`
}

type User struct {
	AlgoAddress string `envconfig:"MY_ALGO_ADDRESS" required:"true"`
}

func NewConfig() {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	envconfig.MustProcess("", &Cfg)
}
