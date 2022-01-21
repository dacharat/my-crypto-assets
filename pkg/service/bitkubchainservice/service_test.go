package bitkubchainservice_test

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/external/web3eth"
	"github.com/dacharat/my-crypto-assets/pkg/external/web3eth/mock_web3eth"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubchainservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("Platform", func(tt *testing.T) {
		tt.Run("return BitkubChain", func(ttt *testing.T) {
			svc, _, finish := newBitkubChainTestSvc(ttt)
			defer finish()

			platform := svc.Platform()

			require.Equal(ttt, platform, shared.BitkubChain)
		})
	})

	t.Run("GetAccount", func(tt *testing.T) {
		ctx := context.Background()

		tt.Run("GetAccountBalance error", func(ttt *testing.T) {
			svc, mockSvc, finish := newBitkubChainTestSvc(ttt)
			defer finish()

			account := common.HexToAddress("0x77eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address1 := common.HexToAddress("0x67eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address2 := common.HexToAddress("0x726613C4494C60B7dCdeA5BE2846180C1DAfBE8B")

			mockSvc.mockWeb3.EXPECT().GetAccountBalance(ctx, account).Return(nil, errors.New("error"))
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address1, account).Return(mockTokenInfo("a"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address2, account).Return(mockTokenInfo("b"), nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: account.Hex(),
			})

			require.Error(ttt, err)
		})

		tt.Run("GetTokenBalance error", func(ttt *testing.T) {
			svc, mockSvc, finish := newBitkubChainTestSvc(ttt)
			defer finish()

			account := common.HexToAddress("0x77eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address1 := common.HexToAddress("0x67eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address2 := common.HexToAddress("0x726613C4494C60B7dCdeA5BE2846180C1DAfBE8B")

			mockSvc.mockWeb3.EXPECT().GetAccountBalance(ctx, account).Return(big.NewInt(1), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address1, account).Return(mockTokenInfo("a"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address2, account).Return(mockTokenInfo("b"), errors.New("error"))

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: account.Hex(),
			})

			require.Error(ttt, err)
		})

		tt.Run("return account", func(ttt *testing.T) {
			svc, mockSvc, finish := newBitkubChainTestSvc(ttt)
			defer finish()

			account := common.HexToAddress("0x77eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address1 := common.HexToAddress("0x67eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address2 := common.HexToAddress("0x726613C4494C60B7dCdeA5BE2846180C1DAfBE8B")

			mockSvc.mockWeb3.EXPECT().GetAccountBalance(ctx, account).Return(big.NewInt(1), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address1, account).Return(mockTokenInfo("a"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address2, account).Return(mockTokenInfo("b"), nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: account.Hex(),
			})

			require.NoError(ttt, err)
		})
	})
}

type bitkubChainTestSvc struct {
	mockWeb3 *mock_web3eth.MockIWeb3Eth
}

func newBitkubChainTestSvc(t gomock.TestHelper) (shared.IAssetsService, bitkubChainTestSvc, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := bitkubChainTestSvc{
		mockWeb3: mock_web3eth.NewMockIWeb3Eth(ctrl),
	}

	svc := bitkubchainservice.NewService(mockSvc.mockWeb3)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}

func mockTokenInfo(symbol string) web3eth.Token {
	return web3eth.Token{
		Balance:  big.NewInt(1),
		Symbol:   symbol,
		Decimals: 18,
	}
}
