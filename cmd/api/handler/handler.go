package handler

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	algoranService algorandservice.IAlgorandService
}

func NewHandler(algo algorandservice.IAlgorandService) Handler {
	return Handler{
		algoranService: algo,
	}
}

func (h Handler) GetAccountBalanceHandler(c *gin.Context) {
	ctx := c.Request.Context()

	res, err := h.algoranService.GetAccount(ctx, config.Cfg.User.AlgoAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": res})
}
