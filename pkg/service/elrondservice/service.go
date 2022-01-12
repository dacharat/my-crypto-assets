package elrondservice

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/dacharat/my-crypto-assets/pkg/external/elrond"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
)

const decimal = 18

type service struct {
	elrondApi elrond.IElrond
}

func NewService(elrondApi elrond.IElrond) shared.IAssetsService {
	return &service{
		elrondApi: elrondApi,
	}
}

func (s *service) Platform() shared.Platform {
	return shared.ElrondChain
}

func (s *service) GetAccount(ctx context.Context, req shared.GetAccountReq) (shared.Account, error) {
	acc, dele, err := s.asyncFetchAccount(ctx, req.WalletAddress)
	if err != nil {
		return shared.Account{}, err
	}

	assets := mapDelegationToAsset(dele)
	assets = append(assets, mapAccountToAsset(acc))

	return shared.Account{
		Platform:     shared.ElrondChain,
		Address:      req.WalletAddress,
		Assets:       assets,
		NeedCgkPrice: true,
	}, nil
}

func (s *service) asyncFetchAccount(ctx context.Context, address string) (elrond.GetAccountResponse, []elrond.GetAccountDelegationResponse, error) {
	maxConcurrent := 2
	var (
		ch          = make(chan error, maxConcurrent)
		account     elrond.GetAccountResponse
		delegations []elrond.GetAccountDelegationResponse
	)

	go func() {
		acc, err := s.elrondApi.GetAccount(ctx, address)
		if err == nil {
			account = acc
		}
		ch <- err
	}()

	go func() {
		dele, err := s.elrondApi.GetAccountDelegation(ctx, address)
		if err == nil {
			delegations = dele
		}
		ch <- err
	}()

	var err error
	for i := 0; i < maxConcurrent; i++ {
		errCh := <-ch
		if errCh != nil {
			err = fmt.Errorf("%d: %w", i, errCh)
		}
	}

	return account, delegations, err
}

func mapAccountToAsset(account elrond.GetAccountResponse) *shared.Asset {
	amount, _ := strconv.ParseInt(account.Balance, 10, 64)
	return &shared.Asset{
		Name:   "EGLD",
		Amount: number.BigIntToFloat(big.NewInt(amount), decimal),
	}
}

func mapDelegationToAsset(delegations []elrond.GetAccountDelegationResponse) shared.Assets {
	assets := make(shared.Assets, len(delegations))
	for i, delegation := range delegations {
		stake, _ := strconv.ParseInt(delegation.UserActiveStake, 10, 64)
		reward, _ := strconv.ParseInt(delegation.ClaimableRewards, 10, 64)
		assets[i] = &shared.Asset{
			Name:   "DELEG-EGLD",
			Amount: number.BigIntToFloat(big.NewInt(stake+reward), decimal),
		}
	}

	return assets
}
