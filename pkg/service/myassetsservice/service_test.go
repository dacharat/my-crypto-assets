package myassetsservice_test

import (
	"context"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko/mock_coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/shared/mock_assets_service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAllAssets", func(tt *testing.T) {
		tt.Run("success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newMyAssetsTestSvc(ttt)
			defer finish()

			mockSvc.mockCoinGecko.EXPECT().GetAllPrice(ctx)
			mockSvc.mockAssetsService.EXPECT().Platform().AnyTimes()
			mockSvc.mockAssetsService.EXPECT().GetAccount(gomock.Any(), shared.GetAccountReq{}).Times(3)

			assets, err := svc.GetAllAssets(ctx)

			require.NoError(ttt, err)
			require.Equal(ttt, len(assets), 3)
		})
	})
}

type myAssetsServiceMock struct {
	mockAssetsService *mock_assets_service.MockIAssetsService
	mockCoinGecko     *mock_coingecko.MockICoingecko
}

func newMyAssetsTestSvc(t gomock.TestReporter) (myassetsservice.IMyAssetsService, myAssetsServiceMock, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.User{
		AlgoAddress: "algo_address",
	}

	mockSvc := myAssetsServiceMock{
		mockAssetsService: mock_assets_service.NewMockIAssetsService(ctrl),
		mockCoinGecko:     mock_coingecko.NewMockICoingecko(ctrl),
	}

	assetsSvcs := []shared.IAssetsService{mockSvc.mockAssetsService, mockSvc.mockAssetsService, mockSvc.mockAssetsService}

	svc := myassetsservice.NewService(assetsSvcs, mockSvc.mockCoinGecko, cfg)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
