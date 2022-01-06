package binanceservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/external/binance"
	"github.com/dacharat/my-crypto-assets/pkg/external/binance/mock_binance"
	"github.com/dacharat/my-crypto-assets/pkg/service/binanceservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAccount", func(tt *testing.T) {
		tt.Run("should return error from get binance account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			mockSvc.mockBinance.EXPECT().GetAccount(ctx).Return(binance.GetAccountResponse{}, errors.New("error"))

			_, err := svc.GetAccount(ctx)
			require.Error(ttt, err)
		})

		tt.Run("should return error from get binance tricker", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			mockSvc.mockBinance.EXPECT().GetAccount(ctx).Return(binance.GetAccountResponse{}, nil)
			mockSvc.mockBinance.EXPECT().GetTricker(ctx).Return(map[string]float64{}, errors.New("error"))

			_, err := svc.GetAccount(ctx)
			require.Error(ttt, err)
		})

		tt.Run("should return binance account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			mockSvc.mockBinance.EXPECT().GetAccount(ctx).Return(binance.GetAccountResponse{}, nil)
			mockSvc.mockBinance.EXPECT().GetTricker(ctx).Return(map[string]float64{}, nil)

			_, err := svc.GetAccount(ctx)
			require.NoError(ttt, err)
		})
	})
}

type binanceServiceMock struct {
	mockBinance *mock_binance.MockIBinance
}

func newBinanceTestSvc(t gomock.TestReporter) (shared.IAssetsService, binanceServiceMock, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := binanceServiceMock{
		mockBinance: mock_binance.NewMockIBinance(ctrl),
	}

	svc := binanceservice.NewService(mockSvc.mockBinance)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
