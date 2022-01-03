package handler

import (
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
	h.lineSvc.SendFlex(c.Request.Context())
	c.JSON(http.StatusOK, gin.H{})
}
