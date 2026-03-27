package wallet

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	domain "github.com/sowens81/go-wallet-microservice/internal/domain/wallet"
)

// ---- Mock Repo ----
type mockRepo struct {
	createFn func(ctx context.Context, w domain.Wallet) error
	getFn    func(ctx context.Context, id, accountID string) (*domain.Wallet, error)
}

func (m *mockRepo) Create(ctx context.Context, w domain.Wallet) error {
	return m.createFn(ctx, w)
}

func (m *mockRepo) GetByID(ctx context.Context, id, accountID string) (*domain.Wallet, error) {
	return m.getFn(ctx, id, accountID)
}

func TestService_CreateWallet(t *testing.T) {
	mock := &mockRepo{
		createFn: func(ctx context.Context, w domain.Wallet) error {
			return nil
		},
	}

	service := NewService(mock)

	err := service.CreateWallet(context.Background(), domain.Wallet{
		ID:        "1",
		AccountID: "acc1",
		Balance:   100,
	})

	assert.NoError(t, err)
}

func TestService_GetWallet(t *testing.T) {
	expected := &domain.Wallet{
		ID:        "1",
		AccountID: "acc1",
		Balance:   100,
	}

	mock := &mockRepo{
		getFn: func(ctx context.Context, id, accountID string) (*domain.Wallet, error) {
			return expected, nil
		},
	}

	service := NewService(mock)

	result, err := service.GetWallet(context.Background(), "1", "acc1")

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
