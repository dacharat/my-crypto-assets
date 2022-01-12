package handler

import (
	"errors"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/service/lineservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/myassetsservice"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	assetsSvc myassetsservice.IMyAssetsService
	lineSvc   lineservice.ILineService
}

func NewHandler(assetsSvc myassetsservice.IMyAssetsService, lineSvc lineservice.ILineService) Handler {
	return Handler{
		assetsSvc: assetsSvc,
		lineSvc:   lineSvc,
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

	event, err := h.lineSvc.ParseRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if len(event) == 0 {
		c.JSON(http.StatusOK, gin.H{})
		return
	}

	token := event[0].ReplyToken
	if !h.lineSvc.IsOwner(event[0].Source.UserID) {
		_ = h.lineSvc.ReplyTextMessage(ctx, token, "Not your assets!!")
		c.JSON(http.StatusSeeOther, gin.H{"error": errors.New("invalid user")})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.lineSvc.PushMessage(c.Request.Context(), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
