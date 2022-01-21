package web3eth

import (
	"context"
	"fmt"
	"math/big"

	"github.com/dacharat/my-crypto-assets/pkg/abi/erctoken"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//go:generate mockgen -source=./service.go -destination=./mock_web3eth/mock_service.go -package=mock_web3eth
type IWeb3Eth interface {
	GetAccountBalance(ctx context.Context, account common.Address) (*big.Int, error)
	GetTokenBalance(tokenAddress, account common.Address) (Token, error)
}

type service struct {
	client *ethclient.Client
}

func NewService(uri string) IWeb3Eth {
	client, err := ethclient.Dial(uri)
	if err != nil {
		panic(err)
	}

	return &service{
		client: client,
	}
}

func (s service) GetAccountBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	return s.client.BalanceAt(ctx, account, nil)
}

func (s service) GetTokenBalance(tokenAddress, account common.Address) (Token, error) {
	instance, err := erctoken.NewToken(tokenAddress, s.client)
	max := 3
	if err != nil {
		return Token{}, err
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
	return Token{
		Balance:  balance,
		Symbol:   symbol,
		Decimals: decimal,
	}, errs
}
