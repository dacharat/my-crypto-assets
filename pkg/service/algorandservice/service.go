package algorandservice

import (
	"context"
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
	"github.com/dacharat/my-crypto-assets/pkg/util/price"
)

type service struct {
	api   algorand.IAlgoland
	price coingecko.ICoingecko
}

func NewService(api algorand.IAlgoland, price coingecko.ICoingecko) shared.IAssetsService {
	return &service{
		api:   api,
		price: price,
	}
}

func (s *service) Type() string {
	return string(shared.Algorand)
}

func (s *service) GetAccount(ctx context.Context) (shared.Account, error) {
	res, price, err := s.asyncGetAccountAndPrice(ctx)
	if err != nil {
		return shared.Account{}, err
	}

	return s.mapToAccount(ctx, res, price), nil
}

func (s *service) asyncGetAccountAndPrice(ctx context.Context) (algorand.Account, coingecko.GetPriceResponse, error) {
	maxConcurrent := 2
	c := make(chan error, maxConcurrent)
	defer close(c)
	var res algorand.Account
	var price coingecko.GetPriceResponse

	go func() {
		priceRes, err := s.price.GetPrice(ctx, coingecko.Algo)
		if err == nil {
			price = priceRes
		}
		c <- err
	}()

	go func() {
		accountAddress := config.Cfg.User.AlgoAddress
		account, err := s.api.GetAlgodAccountByID(ctx, accountAddress)
		if err == nil {
			res = account
		}
		c <- err
	}()

	var err error
	for i := 0; i < maxConcurrent; i++ {
		errChan := <-c
		if errChan != nil {
			err = fmt.Errorf("%d: %w", i, errChan)
		}
	}
	return res, price, err
}

func (s *service) mapToAccount(ctx context.Context, resAcount algorand.Account, priceRes coingecko.GetPriceResponse) shared.Account {
	account := shared.Account{
		Platform: shared.Algorand,
		Address:  resAcount.Address,
	}

	algoAmount := toAmount(resAcount.Amount)
	algoPrice := algoAmount * priceRes.Price(coingecko.AlgoCoinID["ALGO"])
	algo := &shared.Asset{
		Amount:        algoAmount,
		ID:            0,
		Name:          "ALGO",
		Price:         algoPrice,
		FormatedPrice: price.Dollar(algoPrice),
	}

	assets := make(shared.Assets, 0, len(resAcount.Assets))
	for _, asset := range resAcount.Assets {
		if asset.IsFrozen {
			continue
		}

		detail, _ := s.api.GetAssetByID(ctx, asset.AssetID)

		assetAmount := toAmount(asset.Amount)
		p := assetAmount * priceRes.Price(coingecko.AlgoCoinID[detail.Asset.Params.Name])
		assets = append(assets, &shared.Asset{
			Amount:        assetAmount,
			ID:            asset.AssetID,
			Name:          detail.Asset.Params.Name,
			Price:         p,
			FormatedPrice: price.Dollar(p),
		})
	}

	account.Assets = append(shared.Assets{algo}, assets...)
	account.TotalPrice = account.Assets.TotalPrice()

	return account
}

func toAmount(amount int) float64 {
	return number.ToFloat(amount, config.Cfg.AlgorandClient.DefaultDecimal)
}
