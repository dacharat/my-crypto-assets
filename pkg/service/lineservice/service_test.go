package lineservice_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dacharat/my-crypto-assets/pkg/external/line/mock_line"
	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Run("SendFlexMessage", func(tt *testing.T) {
		t.Run("should return error line send flex", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().SendFlexMessage(ctx, "", gomock.Any()).Return(errors.New("error"))

			err := svc.SendFlexMessage(ctx, "", nil)
			require.Error(ttt, err)
		})

		t.Run("should send flex success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().SendFlexMessage(ctx, "", gomock.Any())

			err := svc.SendFlexMessage(ctx, "", nil)
			require.NoError(ttt, err)
		})
	})

	t.Run("ReplyTextMessage", func(tt *testing.T) {
		t.Run("should return error line reply message", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			message := "Hello World!!"

			mockSvc.mockLine.EXPECT().ReplyTextMessage(ctx, "", message).Return(errors.New("error"))

			err := svc.ReplyTextMessage(ctx, "", message)
			require.Error(ttt, err)
		})

		t.Run("should reply message success", func(ttt *testing.T) {
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
		t.Run("should return error line push message", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().PushMessage(ctx, gomock.Any()).Return(errors.New("error"))

			err := svc.PushMessage(ctx, nil)
			require.Error(ttt, err)
		})

		t.Run("should push message success", func(ttt *testing.T) {
			ctx := context.Background()
			svc, mockSvc, finish := newBitkubTestSvc(ttt)
			defer finish()

			mockSvc.mockLine.EXPECT().PushMessage(ctx, gomock.Any())

			err := svc.PushMessage(ctx, nil)
			require.NoError(ttt, err)
		})
	})
}

type lineServiceMock struct {
	mockLine *mock_line.MockILine
}

func newBitkubTestSvc(t gomock.TestReporter) (lineservice.ILineService, lineServiceMock, func()) {
	ctrl := gomock.NewController(t)

	mockSvc := lineServiceMock{
		mockLine: mock_line.NewMockILine(ctrl),
	}

	svc := lineservice.NewService(mockSvc.mockLine)

	finish := func() {
		ctrl.Finish()
	}

	return svc, mockSvc, finish
}
