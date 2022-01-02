package algorandservice

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
	"github.com/dacharat/my-crypto-assets/pkg/util/price"
)

type IAlgorandService interface {
	GetAccount(ctx context.Context, account string) (Account, error)
}

type service struct {
	api   algorand.IAlgoland
	price coingecko.ICoingecko
}

func NewService(api algorand.IAlgoland, price coingecko.ICoingecko) IAlgorandService {
	return &service{
		api:   api,
		price: price,
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
	priceRes, _ := s.price.GetPrice(ctx, coingecko.Algo)

	resAcount := res.Account
	account := Account{
		Address: resAcount.Address,
	}

	algoAmount := toAmount(resAcount.Amount)
	algoPrice := algoAmount * priceRes.Price(coingecko.AlgoCoinID["ALGO"])
	algo := Asset{
		Amount:        algoAmount,
		ID:            0,
		Name:          "ALGO",
		Price:         algoPrice,
		FormatedPrice: price.Dollar(algoPrice),
	}

	assets := make([]Asset, 0, len(resAcount.Assets))
	for _, asset := range resAcount.Assets {
		if asset.IsFrozen {
			continue
		}

		detail, _ := s.api.GetAssetByID(ctx, asset.AssetID)

		assetAmount := toAmount(asset.Amount)
		p := assetAmount * priceRes.Price(coingecko.AlgoCoinID[detail.Asset.Params.Name])
		assets = append(assets, Asset{
			Amount:        assetAmount,
			ID:            asset.AssetID,
			Name:          detail.Asset.Params.Name,
			Price:         p,
			FormatedPrice: price.Dollar(p),
		})
	}

	account.Assets = append([]Asset{algo}, assets...)
	account.TotalPirce = price.Dollar(account.Assets.TotalPrice())

	return account
}

func toAmount(amount int) float64 {
	return number.ToFloat(amount, config.Cfg.AlgorandClient.DefaultDecimal)
}
