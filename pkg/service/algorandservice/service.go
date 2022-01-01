package algorandservice

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
)

type IAlgorandService interface {
	GetAccount(ctx context.Context, account string) (Account, error)
}

type service struct {
	api algorand.IAlgoland
}

func NewService(api algorand.IAlgoland) IAlgorandService {
	return &service{
		api: api,
	}
}

func (s *service) GetAccount(ctx context.Context, account string) (Account, error) {
	res, err := s.api.GetAccountByID(ctx, account)
	if err != nil {
		return Account{}, err
	}

	return s.mapToAccount(ctx, res), nil
}

func (s *service) mapToAccount(ctx context.Context, res algorand.AccountResponse) Account {
	resAcount := res.Account
	account := Account{
		Address: resAcount.Address,
		Amount:  toAmount(resAcount.Amount),
	}

	assets := make([]Asset, 0, len(resAcount.Assets))
	for _, asset := range resAcount.Assets {
		if asset.IsFrozen {
			continue
		}

		detail, _ := s.api.GetAssetByID(ctx, asset.AssetID)

		assets = append(assets, Asset{
			Amount: toAmount(asset.Amount),
			ID:     asset.AssetID,
			Name:   detail.Asset.Params.Name,
		})
	}

	account.Assets = assets

	return account
}

func toAmount(amount int) float64 {
	return number.ToFloat(amount, config.Cfg.AlgorandClient.DefaultDecimal)
}
