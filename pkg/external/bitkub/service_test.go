package bitkub_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/bitkub"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient/mock_client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetWallet", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			bitkubSvc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			uri := "https://bitkub.host.com/wallet"

			mockSvc.mockHttpClient.
				EXPECT().
				Post(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := bitkubSvc.GetWallet(ctx)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			bitkubSvc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			uri := "https://bitkub.host.com/wallet"
			wallet := bitkub.GetWalletResponse{}
			walletStr, _ := json.Marshal(wallet)

			mockSvc.mockHttpClient.
				EXPECT().
				Post(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(walletStr)), nil)

			_, err := bitkubSvc.GetWallet(ctx)

			require.NoError(ttt, err)
		})
	})

	t.Run("GetTricker", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			bitkubSvc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			uri := "https://bitkub.host.com/tricker"

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, nil, gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := bitkubSvc.GetTricker(ctx)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			bitkubSvc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			uri := "https://bitkub.host.com/tricker"
			tricker := bitkub.GetTrickerResponse{}
			trickerStr, _ := json.Marshal(tricker)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, nil, gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(trickerStr)), nil)

			_, err := bitkubSvc.GetTricker(ctx)

			require.NoError(ttt, err)
		})
	})
}

type bitkubSvcMock struct {
	mockHttpClient *mock_client.MockIClient
}

func newBitkubTestSvc(t gomock.TestReporter) (bitkub.IBitkub, bitkubSvcMock, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.Bitkub{
		Host:       "https://bitkub.host.com",
		GetWallet:  "/wallet",
		GetTricker: "/tricker",
		ApiKey:     "api-key",
		ApiSecret:  "api-scret",
	}

	mockSvc := bitkubSvcMock{
		mockHttpClient: mock_client.NewMockIClient(ctrl),
	}

	finish := func() {
		ctrl.Finish()
	}

	bitkubSvc := bitkub.NewBitkubService(mockSvc.mockHttpClient, cfg)

	return bitkubSvc, mockSvc, finish
}

func createHttpResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func generateHeaderMock() http.Header {
	header := http.Header{}
	header.Set("X-Btk-Apikey", "api-key")
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")

	return header
}
