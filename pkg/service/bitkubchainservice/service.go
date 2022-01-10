package bitkubchainservice

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dacharat/my-crypto-assets/pkg/abi/erctoken"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewService(conn *ethclient.Client) shared.IAssetsService {
	addresses := []string{"0x67eBD850304c70d983B2d1b93ea79c7CD6c3F6b5", "0x726613C4494C60B7dCdeA5BE2846180C1DAfBE8B"}
	return &service{
		conn:      conn,
		addresses: addresses,
	}
}

type service struct {
	conn      *ethclient.Client
	addresses []string
}

func (s service) Platform() shared.Platform {
	return shared.BitkubChain
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
		if err != nil {
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
	balance, err := s.conn.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, err
	}

	return &shared.Asset{
		Name:   "KUB",
		Amount: number.BigIntToFloat(balance, 18),
	}, nil
}

func (s *service) getTokenInfo(account common.Address, address string) (*shared.Asset, error) {
	tokenAddress := common.HexToAddress(address)
	instance, err := erctoken.NewToken(tokenAddress, s.conn)
	max := 3
	if err != nil {
		return nil, err
	}

	var (
		balance *big.Int
		symbol  string
		decimal uint8
	)

	ch := make(chan error, max)

	go func() {
		bal, err := instance.BalanceOf(&bind.CallOpts{}, account)
		if err == nil {
			balance = bal
		}
		ch <- err
	}()

	go func() {
		sym, err := instance.Symbol(&bind.CallOpts{})
		if err == nil {
			symbol = sym
		}
		ch <- err
	}()

	go func() {
		decimals, err := instance.Decimals(&bind.CallOpts{})
		if err == nil {
			decimal = decimals
		}
		ch <- err
	}()

	var errs error
	for i := 0; i < max; i++ {
		errCh := <-ch
		if errCh != nil {
			errs = fmt.Errorf("%d: %w", i, errCh)
		}
	}
	return &shared.Asset{
		Name:   symbol,
		Amount: number.BigIntToFloat(balance, int(decimal)),
	}, errs
}
