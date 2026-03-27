package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	app "github.com/sowens81/go-wallet-microservice/internal/application/wallet"
	domain "github.com/sowens81/go-wallet-microservice/internal/domain/wallet"
)

type mockService struct {
	createFn func(ctx context.Context, w domain.Wallet) error
	getFn    func(ctx context.Context, id, accountID string) (*domain.Wallet, error)
}

func (m *mockService) CreateWallet(ctx context.Context, w domain.Wallet) error {
	return m.createFn(ctx, w)
}

func (m *mockService) GetWallet(ctx context.Context, id, accountID string) (*domain.Wallet, error) {
	return m.getFn(ctx, id, accountID)
}

func TestHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockService{
		createFn: func(ctx context.Context, w domain.Wallet) error {
			return nil
		},
	}

	handler := NewWalletHandler((*app.Service)(nil))
	handler.service = (*app.Service)(nil) // override not needed if interface used

	router := gin.Default()
	router.POST("/wallets", handler.Create)

	body := `{"id":"1","accountId":"acc1","balance":100}`

	req := httptest.NewRequest(http.MethodPost, "/wallets", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestHandler_GetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mock := &mockService{
		getFn: func(ctx context.Context, id, accountID string) (*domain.Wallet, error) {
			return &domain.Wallet{
				ID:        id,
				AccountID: accountID,
				Balance:   100,
			}, nil
		},
	}

	handler := &WalletHandler{
		service: (*app.Service)(nil),
	}

	router := gin.Default()
	router.GET("/wallets/:id", handler.GetByID)

	req := httptest.NewRequest(http.MethodGet, "/wallets/1?accountId=acc1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
