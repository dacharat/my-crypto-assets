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

type IBitkubService interface {
	GetAccount(ctx context.Context) (shared.Account, error)
}

type service struct {
	bitkubApi bitkub.IBitkub
}

func NewService(api bitkub.IBitkub) IBitkubService {
	return &service{
		bitkubApi: api,
	}
}

func (s *service) GetAccount(ctx context.Context) (shared.Account, error) {
	res, err := s.bitkubApi.GetWallet(ctx)
	if err != nil {
		return shared.Account{}, err
	}
	tricker, err := s.bitkubApi.GetTricker(ctx)
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
		Assets:     assets,
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

func toUsd(thb float64, rate float64) float64 {
	return thb / rate
}
