package handler_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/dacharat/my-crypto-assets/cmd/api/handler"
	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice/mock_line_service"
	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice/mock_my_assets_service"
	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice/mock_platnetwatch_service"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
	"github.com/dacharat/my-crypto-assets/pkg/util/testutil"
	"github.com/golang/mock/gomock"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	t.Run("GetAccountBalanceHandler", func(tt *testing.T) {
		tt.Run("should return 500", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return(nil, errors.New("error"))

			handler.GetAccountBalanceHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 200", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return([]shared.Account{}, nil)

			handler.GetAccountBalanceHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})
	})

	t.Run("LineCallbackHandler", func(tt *testing.T) {
		tt.Run("should return 400 ParseRequest error", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(nil, errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusBadRequest)
		})

		tt.Run("should return 200 with no action", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(nil, nil)

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})

		tt.Run("should return 303 not owner", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(false)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("message"), nil)
			mockHandler.mockLineSvc.EXPECT().ReplyTextMessage(c.Request.Context(), "reply", "Not your assets!!")

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusSeeOther)
		})

		tt.Run("should return 400 invalid message type", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEventsSticker(), nil)
			mockHandler.mockLineSvc.EXPECT().ReplyTextMessage(c.Request.Context(), "reply", "Not support message type: sticker")

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusBadRequest)
		})

		tt.Run("should return 500 get all assets", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("message"), nil)
			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return(nil, errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 500 get incomes", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Planetwatch"), nil)
			mockHandler.mockPlatnetwatchSvc.EXPECT().GetSummary(gomock.Any()).Return(platnetwatchservice.Summary{}, errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 500 get menu", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Menu"), nil)
			mockHandler.mockLineSvc.EXPECT().SendMenuFlexMessage(c.Request.Context(), "reply").Return(errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 500 get asset by platform", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Algorand"), nil)
			mockHandler.mockAssetsSvc.EXPECT().GetAssetByPlatform(gomock.Any(), shared.Algorand).Return(shared.Account{}, errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 500 send flex", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			accounts := []shared.Account{}

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("message"), nil)
			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return(accounts, nil)
			mockHandler.mockLineSvc.EXPECT().SendFlexMessage(c.Request.Context(), "reply", accounts).Return(errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 500 with send planetwatch incomes", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			summary := platnetwatchservice.Summary{}

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Planetwatch"), nil)
			mockHandler.mockPlatnetwatchSvc.EXPECT().GetSummary(gomock.Any()).Return(summary, nil)
			mockHandler.mockLineSvc.EXPECT().SendPlanetwatchFlexMessage(c.Request.Context(), "reply", summary).Return(errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 200 with send asset by platform fail", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			account := shared.Account{}
			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Algorand"), nil)
			mockHandler.mockAssetsSvc.EXPECT().GetAssetByPlatform(gomock.Any(), shared.Algorand).Return(account, nil)
			mockHandler.mockLineSvc.EXPECT().SendAssetFlexMessage(c.Request.Context(), "reply", account).Return(errors.New("error"))

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 200 with assets", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			accounts := []shared.Account{}

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("message"), nil)
			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return(accounts, nil)
			mockHandler.mockLineSvc.EXPECT().SendFlexMessage(c.Request.Context(), "reply", accounts).Return(nil)

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})

		tt.Run("should return 200 with planetwatch incomes", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			summary := platnetwatchservice.Summary{}

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Planetwatch"), nil)
			mockHandler.mockPlatnetwatchSvc.EXPECT().GetSummary(gomock.Any()).Return(summary, nil)
			mockHandler.mockLineSvc.EXPECT().SendPlanetwatchFlexMessage(c.Request.Context(), "reply", summary).Return(nil)

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})

		tt.Run("should return 200 get menu", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Menu"), nil)
			mockHandler.mockLineSvc.EXPECT().SendMenuFlexMessage(c.Request.Context(), "reply").Return(nil)

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})

		tt.Run("should return 200 get asset by platform", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			account := shared.Account{}
			mockHandler.mockLineSvc.EXPECT().IsOwner("owner").Return(true)
			mockHandler.mockLineSvc.EXPECT().ParseRequest(c.Request).Return(createMockEvents("Algorand"), nil)
			mockHandler.mockAssetsSvc.EXPECT().GetAssetByPlatform(gomock.Any(), shared.Algorand).Return(account, nil)
			mockHandler.mockLineSvc.EXPECT().SendAssetFlexMessage(c.Request.Context(), "reply", account).Return(nil)

			handler.LineCallbackHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})
	})

	t.Run("LinePushMessageHandler", func(tt *testing.T) {
		tt.Run("should return 500 from GetAllAssets error", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return(nil, errors.New("error"))

			handler.LinePushMessageHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 500 from GetAllAssets error", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			accounts := []shared.Account{}

			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return(accounts, nil)
			mockHandler.mockLineSvc.EXPECT().PushMessage(gomock.Any(), accounts).Return(errors.New("error"))

			handler.LinePushMessageHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 200", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			accounts := []shared.Account{}

			mockHandler.mockAssetsSvc.EXPECT().GetAllAssets(gomock.Any()).Return(accounts, nil)
			mockHandler.mockLineSvc.EXPECT().PushMessage(gomock.Any(), accounts).Return(nil)

			handler.LinePushMessageHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})
	})

	t.Run("LinePushMessageByPlatformHandler", func(tt *testing.T) {
		tt.Run("should return 500 from GetAllAssets error", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			mockHandler.mockAssetsSvc.EXPECT().GetAssetByPlatform(gomock.Any(), gomock.Any()).Return(shared.Account{}, errors.New("error"))

			handler.LinePushMessageByPlatformHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 500 from GetAllAssets error", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			account := shared.Account{}

			mockHandler.mockAssetsSvc.EXPECT().GetAssetByPlatform(gomock.Any(), gomock.Any()).Return(account, nil)
			mockHandler.mockLineSvc.EXPECT().PushAssetMessage(gomock.Any(), account).Return(errors.New("error"))

			handler.LinePushMessageByPlatformHandler(c)

			require.Equal(ttt, res.Code, http.StatusInternalServerError)
		})

		tt.Run("should return 200", func(ttt *testing.T) {
			res, c := testutil.NewDefaultContext()

			handler, mockHandler, finish := newHandlerTest(ttt)
			defer finish()

			account := shared.Account{}

			mockHandler.mockAssetsSvc.EXPECT().GetAssetByPlatform(gomock.Any(), gomock.Any()).Return(account, nil)
			mockHandler.mockLineSvc.EXPECT().PushAssetMessage(gomock.Any(), account).Return(nil)

			handler.LinePushMessageByPlatformHandler(c)

			require.Equal(ttt, res.Code, http.StatusOK)
		})
	})
}

type handlerMock struct {
	mockAssetsSvc       *mock_my_assets_service.MockIMyAssetsService
	mockLineSvc         *mock_line_service.MockILineService
	mockPlatnetwatchSvc *mock_platnetwatch_service.MockIPlanetwatchService
}

func newHandlerTest(t gomock.TestHelper) (handler.Handler, handlerMock, func()) {
	ctrl := gomock.NewController(t)

	mockHandler := handlerMock{
		mockAssetsSvc:       mock_my_assets_service.NewMockIMyAssetsService(ctrl),
		mockLineSvc:         mock_line_service.NewMockILineService(ctrl),
		mockPlatnetwatchSvc: mock_platnetwatch_service.NewMockIPlanetwatchService(ctrl),
	}

	finish := func() {
		ctrl.Finish()
	}

	handler := handler.NewHandler(mockHandler.mockAssetsSvc, mockHandler.mockLineSvc, mockHandler.mockPlatnetwatchSvc)

	return handler, mockHandler, finish
}

func createMockEvents(message string) []*linebot.Event {
	return []*linebot.Event{{
		ReplyToken: "reply",
		Source:     &linebot.EventSource{UserID: "owner"},
		Message:    linebot.NewTextMessage(message),
	}}
}

func createMockEventsSticker() []*linebot.Event {
	return []*linebot.Event{{
		ReplyToken: "reply",
		Source:     &linebot.EventSource{UserID: "owner"},
		Message:    linebot.NewStickerMessage("1", "message"),
	}}
}
