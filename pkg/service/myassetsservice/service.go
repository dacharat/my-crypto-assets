package myassetsservice

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/shared"
)

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

func NewHandler(assetSvcs ...shared.IAssetsService) IMyAssetsService {
	return &service{
		assetSvcs: assetSvcs,
	}
}

func (s *service) GetAllAssets(ctx context.Context) ([]shared.Account, error) {
	c := make(chan AccountErr, 3)
	defer close(c)

	for _, assetSvc := range s.assetSvcs {
		go channelFunc(ctx, c, assetSvc.GetAccount)
	}

	var data []shared.Account
	for i := 0; i < 3; i++ {
		result := <-c
		if result.Err != nil {
			return nil, result.Err
		}
		data = append(data, result.Account)
	}

	return data, nil
}

func channelFunc(ctx context.Context, c chan AccountErr, fun func(context.Context) (shared.Account, error)) {
	account, err := fun(ctx)
	c <- AccountErr{
		Account: account,
		Err:     err,
	}
}
