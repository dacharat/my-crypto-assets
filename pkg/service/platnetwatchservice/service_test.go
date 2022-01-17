package platnetwatchservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand/mock_algorand"
	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	ctx := context.Background()
	t.Run("GetIncome", func(tt *testing.T) {
		tt.Run("get income error", func(ttt *testing.T) {
			svc, mockSvc, finish := newPlanetwatchTestSvc(ttt)
			defer finish()

			// opts := []interface{}{algorand.WithLimit(10), algorand.WithAssetID(27165954), algorand.WithCurrencyGreaterThan(0)}
			mockSvc.mockAlgoApi.
				EXPECT().
				GetTransaction(ctx, "123", gomock.Any()).
				Return(algorand.AccountTransactionResponse{}, errors.New("error"))

			_, err := svc.GetIncome(ctx)
			require.Error(ttt, err)
		})

		tt.Run("get income success", func(ttt *testing.T) {
			svc, mockSvc, finish := newPlanetwatchTestSvc(ttt)
			defer finish()

			mockSvc.mockAlgoApi.
				EXPECT().
				GetTransaction(ctx, "123", gomock.Any()).
				Return(createMockAccountTransaction(), nil)

			_, err := svc.GetIncome(ctx)
			require.NoError(ttt, err)
		})
	})
}

type planetwatchTestSvc struct {
	mockAlgoApi *mock_algorand.MockIAlgoland
}

func newPlanetwatchTestSvc(t gomock.TestHelper) (platnetwatchservice.IPlanetwatchService, planetwatchTestSvc, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.Algorand{}

	mockSvc := planetwatchTestSvc{
		mockAlgoApi: mock_algorand.NewMockIAlgoland(ctrl),
	}

	svc := platnetwatchservice.NewService(mockSvc.mockAlgoApi, cfg, "123")

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}

func createMockAccountTransaction() algorand.AccountTransactionResponse {
	return algorand.AccountTransactionResponse{
		Transactions: []algorand.Transaction{
			{
				RoundTime: 1642351770,
			},
		},
	}
}
