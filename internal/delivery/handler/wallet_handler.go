package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sowens81/go-wallet-microservice/internal/application/wallet"
	domain "github.com/sowens81/go-wallet-microservice/internal/domain/wallet"
)

type WalletHandler struct {
	service *wallet.Service
}

func NewWalletHandler(service *wallet.Service) *WalletHandler {
	return &WalletHandler{service: service}
}

func (h *WalletHandler) Create(c *gin.Context) {
	var req domain.Wallet

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := h.service.CreateWallet(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *WalletHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	accountID := c.Query("accountId")

	wallet, err := h.service.GetWallet(c.Request.Context(), id, accountID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, wallet)
}
