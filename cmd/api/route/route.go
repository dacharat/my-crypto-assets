package route

import (
	"net/http"

	"github.com/dacharat/my-crypto-assets/cmd/api/handler"
	"github.com/dacharat/my-crypto-assets/pkg/external/algorand"
	"github.com/dacharat/my-crypto-assets/pkg/service/algorandservice"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	h := handler.NewHandler(algorandservice.NewService(algorand.NewAlgolandService()))
	route.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	route.GET("/test", h.GetAccountBalanceHandler)

	return route
}
