package handler

import (
	"errors"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Handler struct {
	assetsSvc myassetsservice.IMyAssetsService
	lineSvc   lineservice.ILineService

	parseReq func(r *http.Request) ([]*linebot.Event, error)
}

func NewHandler(assetsSvc myassetsservice.IMyAssetsService, lineSvc lineservice.ILineService, parseReq func(r *http.Request) ([]*linebot.Event, error)) Handler {
	return Handler{
		assetsSvc: assetsSvc,
		lineSvc:   lineSvc,
		parseReq:  parseReq,
	}
}

func (h Handler) GetAccountBalanceHandler(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.assetsSvc.GetAllAssets(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
}

func (h Handler) LineCallbackHandler(c *gin.Context) {
	ctx := c.Request.Context()

	event, err := h.parseReq(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	token := event[0].ReplyToken
	if token != config.Cfg.Line.UserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("invalid user")})
		return
	}

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

	c.JSON(http.StatusOK, gin.H{})
}

func (h Handler) LinePushMessageHandler(c *gin.Context) {
	ctx := c.Request.Context()

	data, err := h.assetsSvc.GetAllAssets(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = h.lineSvc.PushMessage(c.Request.Context(), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
