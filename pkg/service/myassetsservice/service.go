package myassetsservice

import (
	"context"
	"sort"

	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
)

//go:generate mockgen -source=./service.go -destination=./mock_my_assets_service/mock_service.go -package=mock_my_assets_service
type IMyAssetsService interface {
	GetAllAssets(ctx context.Context) ([]shared.Account, error)
}

func NewService(assetSvcs []shared.IAssetsService, cgk coingecko.ICoingecko) IMyAssetsService {
	return &service{
		assetSvcs: assetSvcs,
		cgk:       cgk,
	}
}

type service struct {
	assetSvcs []shared.IAssetsService
	cgk       coingecko.ICoingecko
}

type AccountErr struct {
	Account shared.Account
	Err     error
}

func (s *service) GetAllAssets(ctx context.Context) ([]shared.Account, error) {
	c := make(chan AccountErr, len(s.assetSvcs))
	defer close(c)

	req := shared.GetAccountReq{}

	for _, assetSvc := range s.assetSvcs {
		go func(svc shared.IAssetsService) {
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

	// sort account by platform
	sort.Slice(data, func(i, j int) bool {
		return string(data[i].Platform) < string(data[j].Platform)
	})

	return data, nil
}
