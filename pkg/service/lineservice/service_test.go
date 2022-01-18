package lineservice_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/external/line/mock_line"
	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/timeutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("IsOwner", func(tt *testing.T) {
		tt.Run("true", func(ttt *testing.T) {
			svc, _, finish := newBitkubTestSvc(ttt)
			defer finish()

			owner := svc.IsOwner("owner")
			require.True(ttt, owner)
		})

		tt.Run("false", func(ttt *testing.T) {
			svc, _, finish := newBitkubTestSvc(ttt)
			defer finish()

			owner := svc.IsOwner("not owner")
			require.False(ttt, owner)
		})
	})

	t.Run("SendFlexMessage", func(tt *testing.T) {
		tt.Run("should return error line send flex", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().SendFlexMessage(ctx, "", gomock.Any()).Return(errors.New("error"))

			err := svc.SendFlexMessage(ctx, "", createMockAccounts())
			require.Error(ttt, err)
		})

		tt.Run("should send flex success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().SendFlexMessage(ctx, "", gomock.Any())

			err := svc.SendFlexMessage(ctx, "", createMockAccounts())
			require.NoError(ttt, err)
		})
	})

	t.Run("ReplyTextMessage", func(tt *testing.T) {
		tt.Run("should return error line reply message", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			message := "Hello World!!"

			mockSvc.mockLine.EXPECT().ReplyTextMessage(ctx, "", message).Return(errors.New("error"))

			err := svc.ReplyTextMessage(ctx, "", message)
			require.Error(ttt, err)
		})

		tt.Run("should reply message success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			message := "Hello World!!"

			mockSvc.mockLine.EXPECT().ReplyTextMessage(ctx, "", message)

			err := svc.ReplyTextMessage(ctx, "", message)
			require.NoError(ttt, err)
		})
	})

	t.Run("PushMessage", func(tt *testing.T) {
		tt.Run("should return error line push message", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().PushMessage(ctx, gomock.Any()).Return(errors.New("error"))

			err := svc.PushMessage(ctx, createMockAccounts())
			require.Error(ttt, err)
		})

		tt.Run("should push message success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().PushMessage(ctx, gomock.Any())

			err := svc.PushMessage(ctx, createMockAccounts())
			require.NoError(ttt, err)
		})
	})

	t.Run("PushPlanetwatchMessage", func(tt *testing.T) {
		tt.Run("should return error line push message", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().PushMessage(ctx, gomock.Any()).Return(errors.New("error"))

			err := svc.PushPlanetwatchMessage(ctx, createMockSummary())
			require.Error(ttt, err)
		})

		tt.Run("should push message success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().PushMessage(ctx, gomock.Any())

			err := svc.PushPlanetwatchMessage(ctx, createMockSummary())
			require.NoError(ttt, err)
		})
	})

	t.Run("SendPlanetwatchFlexMessage", func(tt *testing.T) {
		tt.Run("should return error line send flex", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().SendFlexMessage(ctx, "", gomock.Any()).Return(errors.New("error"))

			err := svc.SendPlanetwatchFlexMessage(ctx, "", createMockSummary())
			require.Error(ttt, err)
		})

		tt.Run("should send flex success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().SendFlexMessage(ctx, "", gomock.Any())

			err := svc.SendPlanetwatchFlexMessage(ctx, "", createMockSummary())
			require.NoError(ttt, err)
		})
	})
}

type lineServiceMock struct {
	mockLine *mock_line.MockILine
}

func newBitkubTestSvc(t gomock.TestReporter) (lineservice.ILineService, lineServiceMock, func()) {
	ctrl := gomock.NewController(t)

	cfg := &config.User{
		MaxAssetsDisplay: 3,
	}

	mockSvc := lineServiceMock{
		mockLine: mock_line.NewMockILine(ctrl),
	}

	svc := lineservice.NewService(mockSvc.mockLine, cfg, "owner")

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}

func createMockAccounts() []shared.Account {
	return []shared.Account{{
		Platform: shared.Algorand,
		Address:  "algoland_address",
		Assets: shared.Assets{
			{
				ID:     123,
				Amount: 1,
				Name:   "name1",
				Price:  1,
			},
			{
				ID:     456,
				Amount: 1,
				Name:   "name2",
				Price:  1,
			},
			{
				ID:     789,
				Amount: 1,
				Name:   "name3",
				Price:  1,
			},
			{
				ID:     100,
				Amount: 1,
				Name:   "name4",
				Price:  1,
			},
		},
	}}
}

func createMockSummary() platnetwatchservice.Summary {
	return platnetwatchservice.Summary{
		Incomes: []*platnetwatchservice.Income{{
			Date:   time.Date(2022, 1, 16, 23, 32, 00, 00, timeutil.BkkLoc),
			Amount: 1,
		}},
	}
}
