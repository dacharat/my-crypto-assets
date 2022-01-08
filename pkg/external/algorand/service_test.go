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

			config.Cfg.AlgorandClient.UseFreeApi = true

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
}

type algorandSvcMock struct {
	mockHttpClient *mock_client.MockIClient
}

func newAlgorandTestSvc(t gomock.TestReporter) (algorand.IAlgoland, algorandSvcMock, func()) {
	ctrl := gomock.NewController(t)
	config.Cfg.AlgorandClient.AlgodHost = "https://algorand.host.com"
	config.Cfg.AlgorandClient.Host = "https://algorand.free-host.com"
	config.Cfg.AlgorandClient.GetAccountPath = "/account/%s"
	config.Cfg.AlgorandClient.GetAssetPath = "/assets/%d"

	mockSvc := algorandSvcMock{
		mockHttpClient: mock_client.NewMockIClient(ctrl),
	}

	finish := func() {
		config.Cfg.AlgorandClient.UseFreeApi = false
		ctrl.Finish()
	}

	algoSvc := algorand.NewAlgolandService(mockSvc.mockHttpClient)

	return algoSvc, mockSvc, finish
}

func createHttpResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
