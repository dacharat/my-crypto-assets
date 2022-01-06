package myassetsservice_test

import (
	"context"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared/mock_assets_service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAllAssets", func(tt *testing.T) {
		tt.Run("", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newMyAssetsTestSvc(ttt)
			defer finish()

			mockSvc.mockAssetsService.EXPECT().GetAccount(ctx).Times(3)

			assets, err := svc.GetAllAssets(ctx)

			require.NoError(ttt, err)
			require.Equal(ttt, len(assets), 3)
		})
	})
}

type myAssetsServiceMock struct {
	mockAssetsService *mock_assets_service.MockIAssetsService
}

func newMyAssetsTestSvc(t gomock.TestReporter) (myassetsservice.IMyAssetsService, myAssetsServiceMock, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := myAssetsServiceMock{
		mockAssetsService: mock_assets_service.NewMockIAssetsService(ctrl),
	}

	svc := myassetsservice.NewService(mockSvc.mockAssetsService, mockSvc.mockAssetsService, mockSvc.mockAssetsService)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
