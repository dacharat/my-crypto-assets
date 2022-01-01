package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Cfg config

type config struct {
	AlgorandClient Algorand
	User           User
}

type Algorand struct {
	Host           string `envconfig:"ALGORAND_HOST" required:"true"`
	GetAccountPath string `envconfig:"ALGORAND_GET_ACCOUNT_PATH" required:"true"`
	GetAssetPath   string `envconfig:"ALGORAND_GET_ASSET_PATH" required:"true"`
	DefaultDecimal int    `envconfig:"DEFAULT_DECIMAL" required:"true"`
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
