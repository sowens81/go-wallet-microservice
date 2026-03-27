package api

import (
	"github.com/gin-gonic/gin"

	"github.com/sowens81/go-wallet-microservice/internal/delivery/handler"
)

type WalletRoutes struct {
	handler *handler.WalletHandler
}

func NewWalletRoutes(handler *handler.WalletHandler) *WalletRoutes {
	return &WalletRoutes{handler: handler}
}

func (r *WalletRoutes) Register(router *gin.Engine) {
	group := router.Group("/wallets")

	group.POST("", r.handler.Create)
	group.GET("/:id", r.handler.GetByID)
}
