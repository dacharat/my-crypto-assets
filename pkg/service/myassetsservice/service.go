package myassetsservice

import (
	"context"
	"sort"

	"github.com/dacharat/my-crypto-assets/pkg/shared"
)

//go:generate mockgen -source=./service.go -destination=./mock_my_assets_service/mock_service.go -package=mock_my_assets_service
type IMyAssetsService interface {
	GetAllAssets(ctx context.Context) ([]shared.Account, error)
}

type service struct {
	assetSvcs []shared.IAssetsService
}

type AccountErr struct {
	Account shared.Account
	Err     error
}

func NewService(assetSvcs ...shared.IAssetsService) IMyAssetsService {
	return &service{
		assetSvcs: assetSvcs,
	}
}

func (s *service) GetAllAssets(ctx context.Context) ([]shared.Account, error) {
	c := make(chan AccountErr, len(s.assetSvcs))
	defer close(c)

	for _, assetSvc := range s.assetSvcs {
		go channelFunc(ctx, c, assetSvc.GetAccount)
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

func channelFunc(ctx context.Context, c chan AccountErr, fun func(context.Context) (shared.Account, error)) {
	account, err := fun(ctx)
	c <- AccountErr{
		Account: account,
		Err:     err,
	}
}
