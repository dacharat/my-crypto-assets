package elrondservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/external/elrond"
	"github.com/dacharat/my-crypto-assets/pkg/external/elrond/mock_elrond"
	"github.com/dacharat/my-crypto-assets/pkg/service/elrondservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAccount", func(tt *testing.T) {
		tt.Run("should return error from get elrond account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			mockSvc.mockElrond.EXPECT().GetAccount(ctx, "elrond_address").Return(elrond.GetAccountResponse{}, errors.New("error"))
			mockSvc.mockElrond.EXPECT().GetAccountToken(ctx, "elrond_address").Return([]elrond.GetAccountTokenResponse{}, errors.New("error"))
			mockSvc.mockElrond.EXPECT().GetAccountDelegation(ctx, "elrond_address").Return([]elrond.GetAccountDelegationResponse{}, nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: "elrond_address",
			})
			require.Error(ttt, err)
		})

		tt.Run("should return error from get elrond tricker", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			mockSvc.mockElrond.EXPECT().GetAccount(ctx, "elrond_address").Return(elrond.GetAccountResponse{}, nil)
			mockSvc.mockElrond.EXPECT().GetAccountToken(ctx, "elrond_address").Return([]elrond.GetAccountTokenResponse{}, nil)
			mockSvc.mockElrond.EXPECT().GetAccountDelegation(ctx, "elrond_address").Return([]elrond.GetAccountDelegationResponse{}, errors.New("error"))

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: "elrond_address",
			})
			require.Error(ttt, err)
		})

		tt.Run("should return elrond account", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			mockSvc.mockElrond.EXPECT().GetAccount(ctx, "elrond_address").Return(elrond.GetAccountResponse{}, nil)
			mockSvc.mockElrond.EXPECT().GetAccountToken(ctx, "elrond_address").Return([]elrond.GetAccountTokenResponse{}, nil)
			mockSvc.mockElrond.EXPECT().GetAccountDelegation(ctx, "elrond_address").Return([]elrond.GetAccountDelegationResponse{}, nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: "elrond_address",
			})
			require.NoError(ttt, err)
		})
	})
}

type elrondServiceMock struct {
	mockElrond *mock_elrond.MockIElrond
}

func newElrondTestSvc(t gomock.TestReporter) (shared.IAssetsService, elrondServiceMock, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := elrondServiceMock{
		mockElrond: mock_elrond.NewMockIElrond(ctrl),
	}

	svc := elrondservice.NewService(mockSvc.mockElrond)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
