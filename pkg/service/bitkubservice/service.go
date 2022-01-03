package bitkubservice

import (
	"context"
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
)

type IBitkubService interface {
	GetAccount(ctx context.Context) (Account, error)
}

type service struct {
	bitkubApi bitkub.IBitkub
}

func NewService(api bitkub.IBitkub) IBitkubService {
	return &service{
		bitkubApi: api,
	}
}

func (s *service) GetAccount(ctx context.Context) (Account, error) {
	res, err := s.bitkubApi.GetWallet(ctx)
	if err != nil {
		return Account{}, err
	}

	assets := mapToAssets(res.Result)
	tricker, err := s.bitkubApi.GetTricker(ctx)
	if err != nil {
		return Account{}, err
	}

	for _, asset := range assets {
		if asset.Name == "THB" {
			asset.Price = asset.Amount
			continue
		}

		key := fmt.Sprintf("THB_%s", asset.Name)
		t := tricker[key]
		asset.Price = asset.Amount * t.Last
	}

	return Account{
		Assets:     assets,
		TotalPrice: assets.TotalPrice(),
	}, nil
}

func mapToAssets(result map[string]float64) Assets {
	assets := make(Assets, 0, len(result))
	for k, v := range result {
		if v == 0 {
			continue
		}

		assets = append(assets, &Asset{
			Name:   k,
			Amount: v,
		})
	}

	return assets
}
