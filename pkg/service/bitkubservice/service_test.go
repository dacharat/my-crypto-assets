package bitkubservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub/mock_bitkub"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAccount", func(tt *testing.T) {
		tt.Run("should return error from get bitkub account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockBitkub.EXPECT().GetWallet(ctx).Return(bitkub.GetWalletResponse{}, errors.New("error"))
			mockSvc.mockBitkub.EXPECT().GetTricker(ctx).Return(bitkub.GetTrickerResponse{}, nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{})
			require.Error(ttt, err)
		})

		tt.Run("should return error from get bitkub tricker", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockBitkub.EXPECT().GetWallet(ctx).Return(bitkub.GetWalletResponse{}, nil)
			mockSvc.mockBitkub.EXPECT().GetTricker(ctx).Return(bitkub.GetTrickerResponse{}, errors.New("error"))

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{})
			require.Error(ttt, err)
		})

		tt.Run("should return bitkub account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockBitkub.EXPECT().GetWallet(ctx).Return(bitkub.GetWalletResponse{}, nil)
			mockSvc.mockBitkub.EXPECT().GetTricker(ctx).Return(bitkub.GetTrickerResponse{}, nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{})
			require.NoError(ttt, err)
		})
	})
}

type bitkubServiceMock struct {
	mockBitkub *mock_bitkub.MockIBitkub
}

func newBitkubTestSvc(t gomock.TestReporter) (shared.IAssetsService, bitkubServiceMock, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := bitkubServiceMock{
		mockBitkub: mock_bitkub.NewMockIBitkub(ctrl),
	}

	svc := bitkubservice.NewService(mockSvc.mockBitkub)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
