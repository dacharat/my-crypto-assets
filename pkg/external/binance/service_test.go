package binance_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/binance"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient/mock_client"
	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAccount", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			binanceSvc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			timeutil.Now = func() time.Time {
				return time.Date(2022, 1, 1, 1, 1, 1, 1, timeutil.BkkLoc)
			}

			nowMilli := timeutil.Now().UnixMilli()

			uri := fmt.Sprintf("%s%s?timestamp=%d&signature=a949160dee9d4c063525ab83812829cf76421f25dc9d3555cb5a2080edcd3809", config.Cfg.Binance.Host, config.Cfg.Binance.GetAccount, nowMilli)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := binanceSvc.GetAccount(ctx)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			binanceSvc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			timeutil.Now = func() time.Time {
				return time.Date(2022, 1, 1, 1, 1, 1, 1, timeutil.BkkLoc)
			}
			nowMilli := timeutil.Now().UnixMilli()
			uri := fmt.Sprintf("%s%s?timestamp=%d&signature=a949160dee9d4c063525ab83812829cf76421f25dc9d3555cb5a2080edcd3809", config.Cfg.Binance.Host, config.Cfg.Binance.GetAccount, nowMilli)
			account := binance.GetAccountResponse{}
			accountStr, _ := json.Marshal(account)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(accountStr)), nil)

			_, err := binanceSvc.GetAccount(ctx)

			require.NoError(ttt, err)
		})
	})

	t.Run("GetSavingBalance", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			binanceSvc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			timeutil.Now = func() time.Time {
				return time.Date(2022, 1, 1, 1, 1, 1, 1, timeutil.BkkLoc)
			}

			nowMilli := timeutil.Now().UnixMilli()

			uri := fmt.Sprintf("%s%s?timestamp=%d&signature=a949160dee9d4c063525ab83812829cf76421f25dc9d3555cb5a2080edcd3809", config.Cfg.Binance.Host, config.Cfg.Binance.GetSaving, nowMilli)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := binanceSvc.GetSavingBalance(ctx)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			binanceSvc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			timeutil.Now = func() time.Time {
				return time.Date(2022, 1, 1, 1, 1, 1, 1, timeutil.BkkLoc)
			}
			nowMilli := timeutil.Now().UnixMilli()
			uri := fmt.Sprintf("%s%s?timestamp=%d&signature=a949160dee9d4c063525ab83812829cf76421f25dc9d3555cb5a2080edcd3809", config.Cfg.Binance.Host, config.Cfg.Binance.GetSaving, nowMilli)
			savingBalance := binance.GetSavingBalanceResponse{}
			savingBalanceStr, _ := json.Marshal(savingBalance)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(savingBalanceStr)), nil)

			_, err := binanceSvc.GetSavingBalance(ctx)

			require.NoError(ttt, err)
		})
	})

	t.Run("GetTricker", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			binanceSvc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			uri := fmt.Sprintf("%s%s", config.Cfg.Binance.Host, config.Cfg.Binance.GetTricker)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := binanceSvc.GetTricker(ctx)

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			binanceSvc, mockSvc, finish := newBinanceTestSvc(ttt)
			defer finish()

			uri := fmt.Sprintf("%s%s", config.Cfg.Binance.Host, config.Cfg.Binance.GetTricker)
			tricker := binance.GetTrickerResponse{}
			trickerStr, _ := json.Marshal(tricker)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, uri, generateHeaderMock(), gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(trickerStr)), nil)

			_, err := binanceSvc.GetTricker(ctx)

			require.NoError(ttt, err)
		})
	})
}

type binanceSvcMock struct {
	mockHttpClient *mock_client.MockIClient
}

func newBinanceTestSvc(t gomock.TestReporter) (binance.IBinance, binanceSvcMock, func()) {
	ctrl := gomock.NewController(t)
	config.Cfg.Binance.Host = "https://binance.host.com"
	config.Cfg.Binance.GetAccount = "/accounts"
	config.Cfg.Binance.GetTricker = "/tricker"
	config.Cfg.Binance.GetSaving = "/saving"
	config.Cfg.Binance.ApiKey = "api-key"
	config.Cfg.Binance.ApiSecret = "api-scret"

	mockSvc := binanceSvcMock{
		mockHttpClient: mock_client.NewMockIClient(ctrl),
	}

	finish := func() {
		ctrl.Finish()
	}

	binanceSvc := binance.NewBinanceService(mockSvc.mockHttpClient)

	return binanceSvc, mockSvc, finish
}

func createHttpResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

func generateHeaderMock() http.Header {
	header := http.Header{}
	header.Set("X-MBX-APIKEY", "api-key")
	header.Set("Content-Type", "application/json")
	header.Set("Accept", "application/json")

	return header
}
