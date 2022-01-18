package platnetwatchservice

import (
	"context"
	"fmt"
	"time"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/util/number"
	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
)

//go:generate mockgen -source=./service.go -destination=./mock_platnetwatch_service/mock_service.go -package=mock_platnetwatch_service
type IPlanetwatchService interface {
	GetSummary(ctx context.Context) (Summary, error)
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

func (s *service) GetSummary(ctx context.Context) (Summary, error) {
	maxConcurrent := 2
	ch := make(chan error, maxConcurrent)

	var (
		incomes []*Income
		streams []*StreamData
	)

	go func() {
		i, err := s.GetIncome(ctx)
		if err == nil {
			incomes = i
		}
		ch <- err
	}()

	go func() {
		s, err := s.GetLastStreamData(ctx)
		if err == nil {
			streams = s
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

	return Summary{
		Incomes:    incomes,
		StreamData: streams,
	}, err
}

func (s *service) GetIncome(ctx context.Context) ([]*Income, error) {
	txn, err := s.algorandApi.GetTransaction(ctx, s.address, algorand.WithLimit(10), algorand.WithAssetID(27165954), algorand.WithCurrencyGreaterThan(0))
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

func (s *service) GetLastStreamData(ctx context.Context) ([]*StreamData, error) {
	txn, err := s.algorandApi.GetTransaction(ctx, s.address, algorand.WithLimit(6), algorand.WithAssetID(27165954))
	if err != nil {
		return nil, err
	}

	streams := make([]*StreamData, 0, 5)
	for _, t := range txn.Transactions {
		if len(streams) >= 5 {
			break
		}
		// skip reward transactions
		if t.AssetTransferTransaction.Amount > 0 {
			continue
		}

		streams = append(streams, &StreamData{
			Date: time.Unix(int64(t.RoundTime), 0).In(timeutil.BkkLoc),
		})
	}

	return streams, nil
}
