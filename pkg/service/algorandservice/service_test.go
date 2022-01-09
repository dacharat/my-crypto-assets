package algorandservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand/mock_algorand"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko/mock_coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
)

func TestService(t *testing.T) {
	t.Run("GetAccount", func(tt *testing.T) {
		tt.Run("should return error from get algorand account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			config.Cfg.User.AlgoAddress = "123"

			mockSvc.mockAlgorand.EXPECT().GetAlgodAccountByID(ctx, "123").Return(algorand.Account{}, errors.New("error"))
			mockSvc.mockCoinGecko.EXPECT().GetPrice(ctx, coingecko.Algo).Return(coingecko.GetPriceResponse{}, nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{})
			require.Error(ttt, err)
		})

		tt.Run("should return algorand account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			mockSvc.mockAlgorand.EXPECT().GetAlgodAccountByID(ctx, "").Return(algorand.Account{
				Address: "",
			}, nil)
			mockSvc.mockCoinGecko.EXPECT().GetPrice(ctx, coingecko.Algo).Return(coingecko.GetPriceResponse{}, nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{})
			require.NoError(ttt, err)
		})
	})
}

type algorandServiceMock struct {
	mockAlgorand  *mock_algorand.MockIAlgoland
	mockCoinGecko *mock_coingecko.MockICoingecko
}

func newAlgorandTestSvc(t gomock.TestReporter) (shared.IAssetsService, algorandServiceMock, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := algorandServiceMock{
		mockAlgorand:  mock_algorand.NewMockIAlgoland(ctrl),
		mockCoinGecko: mock_coingecko.NewMockICoingecko(ctrl),
	}

	svc := algorandservice.NewService(mockSvc.mockAlgorand, mockSvc.mockCoinGecko)

	finish := func() {
		config.Cfg.User.AlgoAddress = ""
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
