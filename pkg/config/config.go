package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AlgorandClient Algorand
	Coingecko      Coingecko
	Bitkub         Bitkub
	Binance        Binance
	ChainRpc       ChainRpc
	Elrond         Elrond
	Line           Line
	User           User
	DevMode        bool   `envconfig:"DEV_MODE" default:"false"`
	Port           string `envconfig:"PORT" default:"8080"`
}

type Algorand struct {
	Host         string `envconfig:"ALGORAND_HOST" required:"true"`
	GetAssetPath string `envconfig:"ALGORAND_GET_ASSET_PATH" default:"/v2/assets/%d"`

	AlgodHost                  string `envconfig:"ALGORAND_ALGOD_HOST" required:"true"`
	ApiKey                     string `envconfig:"ALGORAND_ALGOD_APILEY"`
	GetAccountPath             string `envconfig:"ALGORAND_GET_ACCOUNT_PATH" default:"/v2/accounts/%s"`
	GetAccountTransactionsPath string `envconfig:"ALGORAND_GET_ACCOUNT_TRANSACTIONS_PATH" default:"/v2/accounts/%s/transactions"`

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

type ChainRpc struct {
	Bitkub string `envconfig:"RPC_BITKUB_CLIENT" default:"https://rpc.bitkubchain.io"`
	Bsc    string `envconfig:"RPC_BSC_CLIENT" default:"https://bsc-dataseed.binance.org"`
}

type Elrond struct {
	Host                  string `envconfig:"ELROND_HOST" required:"true"`
	DelegationHost        string `envconfig:"ELROND_DELEGATION_HOST" required:"true"`
	GetAccount            string `envconfig:"ELROND_GET_ACCOUNT" default:"/accounts/%s"`
	GetAccountTokens      string `envconfig:"ELROND_GET_ACCOUNT_TOKENS" default:"/accounts/%s/tokens"`
	GetAccountDelegations string `envconfig:"ELROND_GET_ACCOUNT_DELEGATIONS" default:"/accounts/%s/delegations"`
	GetAccountNfts        string `envconfig:"ELROND_GET_ACCOUNT_NFTS" default:"/accounts/%s/nfts"`
}

type Line struct {
	UserID             string `envconfig:"LINE_USER_ID" required:"true"`
	ChannelSecret      string `envconfig:"LINE_CHANNEL_SECRET" required:"true"`
	ChannelAccessToken string `envconfig:"LINE_CHANNEL_ACCESS_TOKEN" required:"true"`
}

type User struct {
	AlgoAddress      string `envconfig:"MY_ALGO_ADDRESS" required:"true"`
	BitkubAddress    string `envconfig:"MY_BITKUB_ADDRESS" required:"true"`
	BscAddress       string `envconfig:"MY_BSC_ADDRESS" required:"true"`
	ElrondAddress    string `envconfig:"MY_ELROND_ADDRESS" required:"true"`
	MaxAssetsDisplay int    `envconfig:"MAX_ASSETS_DISPLAY" default:"3"`
}

func NewConfig() *Config {
	if os.Getenv("ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	var cfg Config
	envconfig.MustProcess("", &cfg)

	return &cfg
}
