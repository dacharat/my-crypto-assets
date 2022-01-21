package myassetsservice

import (
	"context"
	"sort"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
)

var platformMapper map[shared.Platform]coingecko.CoingeckoMapper = map[shared.Platform]coingecko.CoingeckoMapper{
	shared.Algorand:    coingecko.AlgoCoinID,
	shared.BitkubChain: coingecko.BitkubCoinID,
	shared.ElrondChain: coingecko.ElrondCoinID,
	shared.BSC:         coingecko.BscCoinID,
}

//go:generate mockgen -source=./service.go -destination=./mock_my_assets_service/mock_service.go -package=mock_my_assets_service
type IMyAssetsService interface {
	GetAllAssets(ctx context.Context) ([]shared.Account, error)
}

func NewService(assetSvcs []shared.IAssetsService, cgk coingecko.ICoingecko, cfg *config.User) IMyAssetsService {
	return &service{
		assetSvcs: assetSvcs,
		cgk:       cgk,
		cfg:       cfg,
	}
}

type service struct {
	assetSvcs []shared.IAssetsService
	cgk       coingecko.ICoingecko
	cfg       *config.User
}

type AccountErr struct {
	Account shared.Account
	Err     error
}

func (s *service) GetAllAssets(ctx context.Context) ([]shared.Account, error) {
	// TODO: goroutine
	gckPrice, err := s.cgk.GetAllPrice(ctx)
	if err != nil {
		return nil, err
	}

	data, err := s.asyncGetAccount(ctx)
	if err != nil {
		return nil, err
	}

	// re assign account
	account := make([]shared.Account, len(data))
	for i, d := range data {
		if !d.NeedCgkPrice {
			account[i] = d
			continue
		}

		for _, asset := range d.Assets {
			price := gckPrice.Price(platformMapper[d.Platform][asset.Name])
			asset.Price = price * asset.Amount
		}

		account[i] = shared.Account{
			Platform:   d.Platform,
			Address:    d.Address,
			Assets:     d.Assets.Sort(),
			TotalPrice: d.Assets.TotalPrice(),
		}
	}

	// sort account by platform
	sort.Slice(account, func(i, j int) bool {
		return string(account[i].Platform) < string(account[j].Platform)
	})

	return account, nil
}

func (s *service) asyncGetAccount(ctx context.Context) ([]shared.Account, error) {
	c := make(chan AccountErr, len(s.assetSvcs))
	defer close(c)

	for _, assetSvc := range s.assetSvcs {
		go func(svc shared.IAssetsService) {
			req := shared.GetAccountReq{}
			switch svc.Platform() {
			case shared.Algorand:
				req.WalletAddress = s.cfg.AlgoAddress
			case shared.BitkubChain:
				req.WalletAddress = s.cfg.BitkubAddress
			case shared.BSC:
				req.WalletAddress = s.cfg.BscAddress
			case shared.ElrondChain:
				req.WalletAddress = s.cfg.ElrondAddress
			}

			account, err := svc.GetAccount(ctx, req)
			c <- AccountErr{
				Account: account,
				Err:     err,
			}
		}(assetSvc)
	}

	var data []shared.Account
	for i := 0; i < len(s.assetSvcs); i++ {
		result := <-c
		if result.Err != nil {
			return nil, result.Err
		}
		data = append(data, result.Account)
	}

	return data, nil
}
