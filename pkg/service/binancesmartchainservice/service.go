package binancesmartchainservice

import (
	"context"
	"fmt"

	"github.com/dacharat/my-crypto-assets/pkg/external/web3eth"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
	"github.com/ethereum/go-ethereum/common"
)

func NewService(web3 web3eth.IWeb3Eth) shared.IAssetsService {
	addresses := []string{"0x42981d0bfbaf196529376ee702f2a9eb9092fcb5", "0x0feadcc3824e7f3c12f40e324a60c23ca51627fc", "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82", "0x070a9867ea49ce7afc4505817204860e823489fe"}
	return &service{
		web3:      web3,
		addresses: addresses,
	}
}

type service struct {
	web3      web3eth.IWeb3Eth
	addresses []string
}

func (s service) Platform() shared.Platform {
	return shared.BSC
}

func (s service) GetAccount(ctx context.Context, req shared.GetAccountReq) (shared.Account, error) {
	account := common.HexToAddress(req.WalletAddress)

	ch := make(chan error, len(s.addresses)+1)
	assets := make(shared.Assets, len(s.addresses)+1)
	go func() {
		asset, err := s.getAccountBalance(ctx, account)
		if err == nil {
			assets[len(s.addresses)] = asset
		}
		ch <- err
	}()

	for i, address := range s.addresses {
		go func(index int, addr string) {
			asset, err := s.getTokenInfo(account, addr)
			if err == nil {
				assets[index] = asset
			}
			ch <- err
		}(i, address)
	}

	var err error
	for i := range assets {
		errCh := <-ch
		if errCh != nil {
			err = fmt.Errorf("%d: %w", i, errCh)
		}
	}

	if err != nil {
		return shared.Account{}, err
	}

	return shared.Account{
		Platform:     s.Platform(),
		Address:      req.WalletAddress,
		Assets:       assets.Sort(),
		TotalPrice:   assets.TotalPrice(),
		NeedCgkPrice: true,
	}, nil
}

func (s *service) getAccountBalance(ctx context.Context, account common.Address) (*shared.Asset, error) {
	balance, err := s.web3.GetAccountBalance(ctx, account)
	if err != nil {
		return nil, err
	}

	return &shared.Asset{
		Name:   "BNB",
		Amount: number.BigIntToFloat(balance, 18),
	}, nil
}

func (s *service) getTokenInfo(account common.Address, address string) (*shared.Asset, error) {
	tokenAddress := common.HexToAddress(address)

	info, err := s.web3.GetTokenBalance(tokenAddress, account)
	if err != nil {
		return nil, err
	}

	return &shared.Asset{
		Name:   info.Symbol,
		Amount: number.BigIntToFloat(info.Balance, int(info.Decimals)),
	}, nil
}
