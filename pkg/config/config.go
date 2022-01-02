package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Cfg config

type config struct {
	AlgorandClient Algorand
	Coingecko      Coingecko
	User           User
}

type Algorand struct {
	Host           string `envconfig:"ALGORAND_HOST" required:"true"`
	GetAccountPath string `envconfig:"ALGORAND_GET_ACCOUNT_PATH" default:"/v2/accounts/%s"`
	GetAssetPath   string `envconfig:"ALGORAND_GET_ASSET_PATH" default:"/v2/assets/%d"`
	DefaultDecimal int    `envconfig:"DEFAULT_DECIMAL" default:"6"`
}

type Coingecko struct {
	Host           string `envconfig:"COINGECKGO_HOST" required:"true"`
	GetSimplePrice string `envconfig:"COINGECKGO_GET_SIMPLE_PRICE" default:"/api/v3/simple/price"`
}

type User struct {
	AlgoAddress string `envconfig:"MY_ALGO_ADDRESS" required:"true"`
}

func NewConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	envconfig.MustProcess("", &Cfg)
}
