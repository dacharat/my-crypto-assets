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
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
)

func TestService(t *testing.T) {
	t.Run("GetAccount", func(tt *testing.T) {
		tt.Run("should return error from get algorand account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			mockSvc.mockAlgorand.EXPECT().GetAlgodAccountByID(ctx, "123").Return(algorand.Account{}, errors.New("error"))

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: "123",
			})
			require.Error(ttt, err)
		})

		tt.Run("should return algorand account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			mockSvc.mockAlgorand.EXPECT().GetAlgodAccountByID(ctx, "").Return(algorand.Account{
				Address: "",
			}, nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{})
			require.NoError(ttt, err)
		})
	})
}

type algorandServiceMock struct {
	mockAlgorand *mock_algorand.MockIAlgoland
}

func newAlgorandTestSvc(t gomock.TestReporter) (shared.IAssetsService, algorandServiceMock, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.Algorand{
		DefaultDecimal: 6,
	}

	mockSvc := algorandServiceMock{
		mockAlgorand: mock_algorand.NewMockIAlgoland(ctrl),
	}

	svc := algorandservice.NewService(mockSvc.mockAlgorand, cfg)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
