package coingecko_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/coingecko"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient/mock_client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetPrice", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			coingeckoSvc, mockSvc, finish := newCoingeckoTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, gomock.Any(), nil).
				Return(nil, errors.New("error"))

			_, err := coingeckoSvc.GetPrice(ctx, coingecko.Algo)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			coingeckoSvc, mockSvc, finish := newCoingeckoTestSvc(ttt)
			defer finish()

			cgk := coingecko.GetPriceResponse{}
			cgkStr, _ := json.Marshal(cgk)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, gomock.Any(), nil).
				Return(createHttpResponse(http.StatusOK, string(cgkStr)), nil)

			_, err := coingeckoSvc.GetPrice(ctx, coingecko.Algo)

			require.NoError(ttt, err)
		})
	})

	t.Run("GetAllPrice", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			coingeckoSvc, mockSvc, finish := newCoingeckoTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, gomock.Any(), nil).
				Return(nil, errors.New("error"))

			_, err := coingeckoSvc.GetAllPrice(ctx)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			coingeckoSvc, mockSvc, finish := newCoingeckoTestSvc(ttt)
			defer finish()

			cgk := coingecko.GetPriceResponse{}
			cgkStr, _ := json.Marshal(cgk)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, gomock.Any(), nil).
				Return(createHttpResponse(http.StatusOK, string(cgkStr)), nil)

			_, err := coingeckoSvc.GetAllPrice(ctx)

			require.NoError(ttt, err)
		})
	})
}

type coingeckoSvcMock struct {
	mockHttpClient *mock_client.MockIClient
}

func newCoingeckoTestSvc(t gomock.TestReporter) (coingecko.ICoingecko, coingeckoSvcMock, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.Coingecko{
		Host:           "https://coingecko.host.com",
		GetSimplePrice: "/getsimpleprice",
	}

	mockSvc := coingeckoSvcMock{
		mockHttpClient: mock_client.NewMockIClient(ctrl),
	}

	finish := func() {
		ctrl.Finish()
	}

	coingeckoSvc := coingecko.NewCoingeckoService(mockSvc.mockHttpClient, cfg)

	return coingeckoSvc, mockSvc, finish
}

func createHttpResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
