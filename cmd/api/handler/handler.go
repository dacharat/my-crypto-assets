package handler

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/binanceservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubservice"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	algoranSvc algorandservice.IAlgorandService
	bitkubSvc  bitkubservice.IBitkubService
	binanceSvc binanceservice.IBinanceService
}

func NewHandler(algo algorandservice.IAlgorandService, bitkubSvc bitkubservice.IBitkubService, binanceSvc binanceservice.IBinanceService) Handler {
	return Handler{
		algoranSvc: algo,
		bitkubSvc:  bitkubSvc,
		binanceSvc: binanceSvc,
	}
}

func (h Handler) GetAccountBalanceHandler(c *gin.Context) {
	ctx := c.Request.Context()

	binance, err := h.binanceSvc.GetAccount(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	bitkub, err := h.bitkubSvc.GetAccount(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	algo, err := h.algoranSvc.GetAccount(ctx, config.Cfg.User.AlgoAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "algo": algo, "bitkub": bitkub, "binance": binance})
}
