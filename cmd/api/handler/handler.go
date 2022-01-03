package handler

import (
	"context"
	"net/http"

	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/binanceservice"
	"github.com/dacharat/my-crypto-assets/pkg/service/bitkubservice"
	"github.com/dacharat/my-crypto-assets/pkg/shared"
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

	wow := make(chan AccountErr, 3)
	defer close(wow)
	go channelFunc(ctx, wow, h.binanceSvc.GetAccount)
	go channelFunc(ctx, wow, h.bitkubSvc.GetAccount)
	go channelFunc(ctx, wow, h.algoranSvc.GetAccount)

	var data []shared.Account
	for i := 0; i < 3; i++ {
		result := <-wow
		if result.Err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Err})
			return

		}
		data = append(data, result.Account)
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": data})
}

type AccountErr struct {
	Account shared.Account
	Err     error
}

func channelFunc(ctx context.Context, c chan AccountErr, fun func(context.Context) (shared.Account, error)) {
	account, err := fun(ctx)
	c <- AccountErr{
		Account: account,
		Err:     err,
	}
}
