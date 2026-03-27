package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/sowens81/go-wallet-microservice/internal/delivery/handler"
)

type mockHandler struct{}

func (m *mockHandler) Create(c *gin.Context) {
	c.Status(http.StatusCreated)
}

func (m *mockHandler) GetByID(c *gin.Context) {
	c.Status(http.StatusOK)
}

func TestWalletRoutes_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	h := &handler.WalletHandler{}
	routes := NewWalletRoutes(h)

	routes.Register(router)

	req := httptest.NewRequest(http.MethodGet, "/wallets/1?accountId=acc1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.NotEqual(t, http.StatusNotFound, w.Code)
}
