package bitkubchainservice

import (
	"context"

	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewService(conn *ethclient.Client) shared.IAssetsService {
	return &service{
		conn: conn,
	}
}

type service struct {
	conn *ethclient.Client
}

func (s service) GetAccount(ctx context.Context, req shared.GetAccountReq) (shared.Account, error) {
	return shared.Account{}, nil
}

func (s service) Type() string {
	return string(shared.BitkubChain)
}
