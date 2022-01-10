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

type service struct {
	binanceApi binance.IBinance
}

func NewService(api binance.IBinance) shared.IAssetsService {
	return &service{
		binanceApi: api,
	}
}

func (s *service) Platform() shared.Platform {
	return shared.Binance
}

func (s *service) GetAccount(ctx context.Context, req shared.GetAccountReq) (shared.Account, error) {
	res, tricker, err := s.asyncGetAccountAndTricker(ctx)
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
		Assets:     assets.Sort(),
		TotalPrice: assets.TotalPrice(),
	}, nil
}

func (s *service) asyncGetAccountAndTricker(ctx context.Context) (binance.GetAccountResponse, map[string]float64, error) {
	maxConcurrent := 2
	var (
		ch      = make(chan error, maxConcurrent)
		res     binance.GetAccountResponse
		tricker map[string]float64
	)
	defer close(ch)

	go func() {
		account, err := s.binanceApi.GetAccount(ctx)
		if err == nil {
			res = account
		}
		ch <- err
	}()

	go func() {
		trickerRes, err := s.binanceApi.GetTricker(ctx)
		if err == nil {
			tricker = trickerRes
		}
		ch <- err
	}()

	var err error
	for i := 0; i < maxConcurrent; i++ {
		errChan := <-ch
		if errChan != nil {
			err = fmt.Errorf("%d: %w", i, errChan)
		}
	}

	return res, tricker, err
}
