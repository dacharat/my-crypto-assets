package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/platnetwatchservice"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Handler struct {
	assetsSvc       myassetsservice.IMyAssetsService
	lineSvc         lineservice.ILineService
	platnetwatchSvc platnetwatchservice.IPlanetwatchService
}

func NewHandler(assetsSvc myassetsservice.IMyAssetsService, lineSvc lineservice.ILineService, platnetwatchSvc platnetwatchservice.IPlanetwatchService) Handler {
	return Handler{
		assetsSvc:       assetsSvc,
		lineSvc:         lineSvc,
		platnetwatchSvc: platnetwatchSvc,
	}
}

func (h Handler) GetAccountBalanceHandler(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.assetsSvc.GetAllAssets(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// data, err := h.platnetwatchSvc.GetIncome(ctx)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
}

func (h Handler) LineCallbackHandler(c *gin.Context) {
	ctx := c.Request.Context()

	event, err := h.lineSvc.ParseRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if len(event) == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	e := event[0]
	token := e.ReplyToken
	if !h.lineSvc.IsOwner(e.Source.UserID) {
		_ = h.lineSvc.ReplyTextMessage(ctx, token, "Not your assets!!")
		c.JSON(http.StatusSeeOther, gin.H{"error": errors.New("invalid user")})
		return
	}

	message, ok := e.Message.(*linebot.TextMessage)
	if !ok {
		_ = h.lineSvc.ReplyTextMessage(ctx, token, fmt.Sprintf("Not support message type: %s", e.Message.Type()))
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("cannot cast message type")})
		return
	}

	switch message.Text {
	case "Planetwatch":
		incomes, err := h.platnetwatchSvc.GetIncome(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		err = h.lineSvc.SendPlanetwatchFlexMessage(ctx, token, incomes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	default:
		data, err := h.assetsSvc.GetAllAssets(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		err = h.lineSvc.SendFlexMessage(c.Request.Context(), token, data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) LinePushMessageHandler(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.assetsSvc.GetAllAssets(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.lineSvc.PushMessage(c.Request.Context(), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// incomes, err := h.platnetwatchSvc.GetIncome(ctx)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	// err = h.lineSvc.PushPlanetwatchMessage(ctx, incomes)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{})
}
