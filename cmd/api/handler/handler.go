package handler

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/config"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubservice"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	algoranSvc algorandservice.IAlgorandService
	bitkubSvc  bitkubservice.IBitkubService
}

func NewHandler(algo algorandservice.IAlgorandService, bitkubSvc bitkubservice.IBitkubService) Handler {
	return Handler{
		algoranSvc: algo,
		bitkubSvc:  bitkubSvc,
	}
}

func (h Handler) GetAccountBalanceHandler(c *gin.Context) {
	ctx := c.Request.Context()

	bitkub, err := h.bitkubSvc.GetWallet(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	algo, err := h.algoranSvc.GetAccount(ctx, config.Cfg.User.AlgoAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "algo": algo, "bitkub": bitkub})
}
