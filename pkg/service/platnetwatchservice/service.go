package platnetwatchservice

import (
	"context"
	"time"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
)

//go:generate mockgen -source=./service.go -destination=./mock_platnetwatch_service/mock_service.go -package=mock_platnetwatch_service
type IPlanetwatchService interface {
	GetIncome(ctx context.Context) ([]*Income, error)
}

type service struct {
	algorandApi algorand.IAlgoland
	cfg         *config.Algorand
	address     string
}

func NewService(algorandApi algorand.IAlgoland, cfg *config.Algorand, address string) IPlanetwatchService {
	return &service{
		algorandApi: algorandApi,
		cfg:         cfg,
		address:     address,
	}
}

func (s *service) GetIncome(ctx context.Context) ([]*Income, error) {
	txn, err := s.algorandApi.GetTransaction(ctx, s.address)
	if err != nil {
		return nil, err
	}

	incomes := make([]*Income, len(txn.Transactions))
	for i, t := range txn.Transactions {
		incomes[i] = &Income{
			Date:   time.Unix(int64(t.RoundTime), 0).In(timeutil.BkkLoc),
			Amount: number.ToFloat(t.AssetTransferTransaction.Amount, s.cfg.DefaultDecimal),
		}
	}

	return incomes, nil
}
