package binanceservice

import (
	"context"
	"fmt"
	"strconv"

	"github.com/dacharat/my-crypto-assets/pkg/external/binance"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/stringutil"
)

var (
	stableCoin   = []string{"BUSD", "USDT", "USDC"} // stable coin
	lockedPrefix = "LD"                             // saving account token will start with LD(lending)
)

type IBinanceService interface {
	GetAccount(ctx context.Context) (shared.Account, error)
}

type service struct {
	binanceApi binance.IBinance
}

func NewService(api binance.IBinance) IBinanceService {
	return &service{
		binanceApi: api,
	}
}

func (s *service) GetAccount(ctx context.Context) (shared.Account, error) {
	res, err := s.binanceApi.GetAccount(ctx)
	if err != nil {
		return shared.Account{}, err
	}
	tricker, err := s.binanceApi.GetTricker(ctx)
	if err != nil {
		return shared.Account{}, err
	}

	assets := make(shared.Assets, 0, len(res.Balances))
	for _, balance := range res.Balances {
		free, err := strconv.ParseFloat(balance.Free, 64)
		if err != nil {
			continue
		}

		locked, err := strconv.ParseFloat(balance.Locked, 64)
		if err != nil {
			continue
		}

		amount := free + locked
		if amount <= 0 {
			continue
		}

		var price float64
		if stringutil.ContainStr(balance.Asset, stableCoin) {
			price = amount
		} else if len(balance.Asset) > 2 && balance.Asset[:2] == lockedPrefix {
			price = amount * tricker[fmt.Sprintf("%sUSDT", balance.Asset[2:])]
		} else {
			price = amount * tricker[fmt.Sprintf("%sUSDT", balance.Asset)]
		}

		assets = append(assets, &shared.Asset{
			Name:   balance.Asset,
			Amount: amount,
			Price:  price,
		})
	}

	return shared.Account{
		Platform:   shared.Binance,
		Assets:     assets,
		TotalPrice: assets.TotalPrice(),
	}, nil
}
