package binancesmartchainservice_test

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/external/web3eth"
	"github.com/dacharat/my-crypto-assets/pkg/external/web3eth/mock_web3eth"
	"github.com/dacharat/my-crypto-assets/pkg/service/binancesmartchainservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("Platform", func(tt *testing.T) {
		tt.Run("return BinanceSmartChain", func(ttt *testing.T) {
			svc, _, finish := newBscTestSvc(ttt)
			defer finish()

			platform := svc.Platform()

			require.Equal(ttt, platform, shared.BSC)
		})
	})

	t.Run("GetAccount", func(tt *testing.T) {
		ctx := context.Background()

		tt.Run("GetAccountBalance error", func(ttt *testing.T) {
			svc, mockSvc, finish := newBscTestSvc(ttt)
			defer finish()

			account := common.HexToAddress("0x77eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address1 := common.HexToAddress("0x42981d0bfbaf196529376ee702f2a9eb9092fcb5")
			address2 := common.HexToAddress("0x0feadcc3824e7f3c12f40e324a60c23ca51627fc")
			address3 := common.HexToAddress("0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82")
			address4 := common.HexToAddress("0x070a9867ea49ce7afc4505817204860e823489fe")

			mockSvc.mockWeb3.EXPECT().GetAccountBalance(ctx, account).Return(nil, errors.New("error"))
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address1, account).Return(mockTokenInfo("a"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address2, account).Return(mockTokenInfo("b"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address3, account).Return(mockTokenInfo("c"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address4, account).Return(mockTokenInfo("d"), nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: account.Hex(),
			})

			require.Error(ttt, err)
		})

		tt.Run("GetTokenBalance error", func(ttt *testing.T) {
			svc, mockSvc, finish := newBscTestSvc(ttt)
			defer finish()

			account := common.HexToAddress("0x77eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address1 := common.HexToAddress("0x42981d0bfbaf196529376ee702f2a9eb9092fcb5")
			address2 := common.HexToAddress("0x0feadcc3824e7f3c12f40e324a60c23ca51627fc")
			address3 := common.HexToAddress("0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82")
			address4 := common.HexToAddress("0x070a9867ea49ce7afc4505817204860e823489fe")

			mockSvc.mockWeb3.EXPECT().GetAccountBalance(ctx, account).Return(big.NewInt(1), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address1, account).Return(mockTokenInfo("a"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address2, account).Return(mockTokenInfo("b"), errors.New("error"))
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address3, account).Return(mockTokenInfo("c"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address4, account).Return(mockTokenInfo("d"), nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: account.Hex(),
			})

			require.Error(ttt, err)
		})

		tt.Run("return account", func(ttt *testing.T) {
			svc, mockSvc, finish := newBscTestSvc(ttt)
			defer finish()

			account := common.HexToAddress("0x77eBD850304c70d983B2d1b93ea79c7CD6c3F6b5")
			address1 := common.HexToAddress("0x42981d0bfbaf196529376ee702f2a9eb9092fcb5")
			address2 := common.HexToAddress("0x0feadcc3824e7f3c12f40e324a60c23ca51627fc")
			address3 := common.HexToAddress("0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82")
			address4 := common.HexToAddress("0x070a9867ea49ce7afc4505817204860e823489fe")

			mockSvc.mockWeb3.EXPECT().GetAccountBalance(ctx, account).Return(big.NewInt(1), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address1, account).Return(mockTokenInfo("a"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address2, account).Return(mockTokenInfo("b"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address3, account).Return(mockTokenInfo("c"), nil)
			mockSvc.mockWeb3.EXPECT().GetTokenBalance(address4, account).Return(mockTokenInfo("d"), nil)

			_, err := svc.GetAccount(ctx, shared.GetAccountReq{
				WalletAddress: account.Hex(),
			})

			require.NoError(ttt, err)
		})
	})
}

type bscTestSvc struct {
	mockWeb3 *mock_web3eth.MockIWeb3Eth
}

func newBscTestSvc(t gomock.TestHelper) (shared.IAssetsService, bscTestSvc, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := bscTestSvc{
		mockWeb3: mock_web3eth.NewMockIWeb3Eth(ctrl),
	}

	svc := binancesmartchainservice.NewService(mockSvc.mockWeb3)

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
