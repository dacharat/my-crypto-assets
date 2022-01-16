package algorand_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient/mock_client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAlgodAccountByID", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			algoSvc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://algorand.host.com/account/123", http.Header{}).
				Return(nil, errors.New("error"))

			_, err := algoSvc.GetAlgodAccountByID(ctx, "123")

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			algoSvc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			account := algorand.Account{}
			accountStr, _ := json.Marshal(account)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://algorand.host.com/account/123", http.Header{}).
				Return(createHttpResponse(http.StatusOK, string(accountStr)), nil)

			_, err := algoSvc.GetAlgodAccountByID(ctx, "123")

			require.NoError(ttt, err)
		})

		tt.Run("should get success by free API", func(ttt *testing.T) {
			ctx := context.Background()
			algoSvc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			mockSvc.mockConfig.UseFreeApi = true

			account := algorand.AccountResponse{}
			accountStr, _ := json.Marshal(account)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://algorand.free-host.com/account/123", nil).
				Return(createHttpResponse(http.StatusOK, string(accountStr)), nil)

			_, err := algoSvc.GetAlgodAccountByID(ctx, "123")

			require.NoError(ttt, err)
		})
	})

	t.Run("GetAssetByID", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			algoSvc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://algorand.free-host.com/assets/123", nil).
				Return(nil, errors.New("error"))

			_, err := algoSvc.GetAssetByID(ctx, 123)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			algoSvc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			assets := algorand.AssetResponse{}
			assetsStr, _ := json.Marshal(assets)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://algorand.free-host.com/assets/123", nil).
				Return(createHttpResponse(http.StatusOK, string(assetsStr)), nil)

			_, err := algoSvc.GetAssetByID(ctx, 123)

			require.NoError(ttt, err)
		})
	})

	t.Run("GetTransaction", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			algoSvc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://algorand.free-host.com/account/123/transactions?limit=10&asset-id=27165954&currency-greater-than=0", nil, gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := algoSvc.GetTransaction(ctx, "123")

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			algoSvc, mockSvc, finish := newAlgorandTestSvc(ttt)
			defer finish()

			assets := algorand.AccountTransactionResponse{}
			assetsStr, _ := json.Marshal(assets)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://algorand.free-host.com/account/123/transactions?limit=10&asset-id=27165954&currency-greater-than=0", nil, gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(assetsStr)), nil)

			_, err := algoSvc.GetTransaction(ctx, "123")

			require.NoError(ttt, err)
		})
	})
}

type algorandSvcMock struct {
	mockHttpClient *mock_client.MockIClient
	mockConfig     *config.Algorand
}

func newAlgorandTestSvc(t gomock.TestReporter) (algorand.IAlgoland, algorandSvcMock, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.Algorand{
		Host:                       "https://algorand.free-host.com",
		AlgodHost:                  "https://algorand.host.com",
		GetAccountPath:             "/account/%s",
		GetAssetPath:               "/assets/%d",
		GetAccountTransactionsPath: "/account/%s/transactions",
	}

	mockSvc := algorandSvcMock{
		mockHttpClient: mock_client.NewMockIClient(ctrl),
		mockConfig:     cfg,
	}

	finish := func() {
		cfg.UseFreeApi = false
		ctrl.Finish()
	}

	algoSvc := algorand.NewAlgolandService(mockSvc.mockHttpClient, cfg)

	return algoSvc, mockSvc, finish
}

func createHttpResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
