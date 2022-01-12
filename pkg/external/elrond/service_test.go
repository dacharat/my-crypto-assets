package elrond_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/elrond"
	"github.com/dacharat/my-crypto-assets/pkg/util/httpclient/mock_client"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("GetAccount", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			elrondSvc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://elrond.host.com/account/elrond_address", nil, gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := elrondSvc.GetAccount(ctx, "elrond_address")

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			coingeckoSvc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			cgk := elrond.GetAccountResponse{}
			cgkStr, _ := json.Marshal(cgk)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://elrond.host.com/account/elrond_address", nil, gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(cgkStr)), nil)

			_, err := coingeckoSvc.GetAccount(ctx, "elrond_address")

			require.NoError(ttt, err)
		})
	})

	t.Run("GetAccountDelegation", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			elrondSvc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://delegations.elrond.host.com/account/elrond_address/delagations", nil, gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := elrondSvc.GetAccountDelegation(ctx, "elrond_address")

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			coingeckoSvc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			cgk := []elrond.GetAccountDelegationResponse{}
			cgkStr, _ := json.Marshal(cgk)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://delegations.elrond.host.com/account/elrond_address/delagations", nil, gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(cgkStr)), nil)

			_, err := coingeckoSvc.GetAccountDelegation(ctx, "elrond_address")

			require.NoError(ttt, err)
		})
	})

	t.Run("GetAccountNfts", func(tt *testing.T) {
		tt.Run("should get error", func(ttt *testing.T) {
			ctx := context.Background()
			elrondSvc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://elrond.host.com/account/elrond_address/nfts?type=MetaESDT", nil, gomock.Any()).
				Return(nil, errors.New("error"))

			_, err := elrondSvc.GetAccountNfts(ctx, "elrond_address")

			require.Error(ttt, err)
		})

		tt.Run("should get success", func(ttt *testing.T) {
			ctx := context.Background()
			coingeckoSvc, mockSvc, finish := newElrondTestSvc(ttt)
			defer finish()

			cgk := []elrond.GetAccountNftResponse{}
			cgkStr, _ := json.Marshal(cgk)

			mockSvc.mockHttpClient.
				EXPECT().
				Get(ctx, "https://elrond.host.com/account/elrond_address/nfts?type=MetaESDT", nil, gomock.Any()).
				Return(createHttpResponse(http.StatusOK, string(cgkStr)), nil)

			_, err := coingeckoSvc.GetAccountNfts(ctx, "elrond_address")

			require.NoError(ttt, err)
		})
	})
}

type elrondSvcMock struct {
	mockHttpClient *mock_client.MockIClient
}

func newElrondTestSvc(t gomock.TestReporter) (elrond.IElrond, elrondSvcMock, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.Elrond{
		Host:                  "https://elrond.host.com",
		DelegationHost:        "https://delegations.elrond.host.com",
		GetAccount:            "/account/%s",
		GetAccountDelegations: "/account/%s/delagations",
		GetAccountNfts:        "/account/%s/nfts",
	}

	mockSvc := elrondSvcMock{
		mockHttpClient: mock_client.NewMockIClient(ctrl),
	}

	finish := func() {
		ctrl.Finish()
	}

	elrondSvc := elrond.NewService(mockSvc.mockHttpClient, cfg)

	return elrondSvc, mockSvc, finish
}

func createHttpResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
