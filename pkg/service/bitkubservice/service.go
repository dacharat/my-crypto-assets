package bitkubservice

import (
	"context"
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
)

const (
	thbToUsdPair = "THB_USDT"
	thbCurrency  = "THB"
	tokenPair    = "THB_%s"
)

type service struct {
	bitkubApi bitkub.IBitkub
}

func NewService(api bitkub.IBitkub) shared.IAssetsService {
	return &service{
		bitkubApi: api,
	}
}

func (s *service) Platform() shared.Platform {
	return shared.Bitkub
}

func (s *service) GetAccount(ctx context.Context, req shared.GetAccountReq) (shared.Account, error) {
	res, tricker, err := s.asyncGetWalletAndTricker(ctx)
	if err != nil {
		return shared.Account{}, err
	}

	usdRate := tricker[thbToUsdPair].Last

	assets := mapToAssets(res.Result)
	for _, asset := range assets {
		if asset.Name == thbCurrency {
			asset.Price = toUsd(asset.Amount, usdRate)
			continue
		}

		key := fmt.Sprintf(tokenPair, asset.Name)
		t := tricker[key]
		asset.Price = toUsd(asset.Amount*t.Last, usdRate)
	}

	return shared.Account{
		Platform:   shared.Bitkub,
		Assets:     assets.Sort(),
		TotalPrice: assets.TotalPrice(),
	}, nil
}

func mapToAssets(result map[string]float64) shared.Assets {
	assets := make(shared.Assets, 0, len(result))
	for k, v := range result {
		if v == 0 {
			continue
		}

		assets = append(assets, &shared.Asset{
			Name:   k,
			Amount: v,
		})
	}

	return assets
}

func (s *service) asyncGetWalletAndTricker(ctx context.Context) (bitkub.GetWalletResponse, bitkub.GetTrickerResponse, error) {
	maxConcurrent := 2
	var (
		ch      = make(chan error, maxConcurrent)
		res     bitkub.GetWalletResponse
		tricker bitkub.GetTrickerResponse
	)
	defer close(ch)

	go func() {
		wallet, err := s.bitkubApi.GetWallet(ctx)
		if err == nil {
			res = wallet
		}
		ch <- err
	}()

	go func() {
		trickerRes, err := s.bitkubApi.GetTricker(ctx)
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

func toUsd(thb float64, rate float64) float64 {
	return thb / rate
}
