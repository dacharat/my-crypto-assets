package myassetsservice

import (
	"context"
	"errors"
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
	GetAssetByPlatform(ctx context.Context, platform shared.Platform) (shared.Account, error)
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

	// re assign accounts
	accounts := make([]shared.Account, len(data))
	for i, d := range data {
		if !d.NeedCgkPrice {
			accounts[i] = d
			continue
		}

		accounts[i] = mapAccountAndPrice(d, gckPrice)
	}

	// sort account by platform
	sort.Slice(accounts, func(i, j int) bool {
		return string(accounts[i].Platform) < string(accounts[j].Platform)
	})

	return accounts, nil
}

func (s *service) GetAssetByPlatform(ctx context.Context, platform shared.Platform) (shared.Account, error) {
	svc := s.findAssetSvc(platform)
	if svc == nil {
		return shared.Account{}, errors.New("not found asset service")
	}

	account, err := svc.GetAccount(ctx, s.createGetAccountReq(platform))
	if err != nil {
		return account, err
	}

	if !account.NeedCgkPrice {
		return account, err
	}

	gckPrice, err := s.cgk.GetAllPrice(ctx)
	if err != nil {
		return account, err
	}

	return mapAccountAndPrice(account, gckPrice), nil
}

func (s *service) asyncGetAccount(ctx context.Context) ([]shared.Account, error) {
	c := make(chan AccountErr, len(s.assetSvcs))
	defer close(c)

	for _, assetSvc := range s.assetSvcs {
		go func(svc shared.IAssetsService) {
			req := s.createGetAccountReq(svc.Platform())

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

func (s *service) findAssetSvc(platform shared.Platform) shared.IAssetsService {
	for _, svc := range s.assetSvcs {
		if svc.Platform() == platform {
			return svc
		}
	}

	return nil
}

func (s *service) createGetAccountReq(platform shared.Platform) shared.GetAccountReq {
	req := shared.GetAccountReq{}
	switch platform {
	case shared.Algorand:
		req.WalletAddress = s.cfg.AlgoAddress
	case shared.BitkubChain:
		req.WalletAddress = s.cfg.BitkubAddress
	case shared.BSC:
		req.WalletAddress = s.cfg.BscAddress
	case shared.ElrondChain:
		req.WalletAddress = s.cfg.ElrondAddress
	}

	return req
}

func mapAccountAndPrice(account shared.Account, gckPrice coingecko.GetPriceResponse) shared.Account {
	for _, asset := range account.Assets {
		price := gckPrice.Price(platformMapper[account.Platform][asset.Name])
		asset.Price = price * asset.Amount
	}

	return shared.Account{
		Platform:   account.Platform,
		Address:    account.Address,
		Assets:     account.Assets.Sort(),
		TotalPrice: account.Assets.TotalPrice(),
	}
}
