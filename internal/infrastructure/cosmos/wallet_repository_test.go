package cosmos

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/stretchr/testify/assert"

	"github.com/sowens81/go-wallet-microservice/internal/domain/wallet"
)

type mockContainer struct {
	createFn func(ctx context.Context, pk azcosmos.PartitionKey, body []byte, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error)
	readFn   func(ctx context.Context, pk azcosmos.PartitionKey, id string, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error)
}

func (m *mockContainer) CreateItem(ctx context.Context, pk azcosmos.PartitionKey, body []byte, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error) {
	return m.createFn(ctx, pk, body, o)
}

func (m *mockContainer) ReadItem(ctx context.Context, pk azcosmos.PartitionKey, id string, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error) {
	return m.readFn(ctx, pk, id, o)
}

func TestWalletRepository_Create(t *testing.T) {
	mock := &mockContainer{
		createFn: func(ctx context.Context, pk azcosmos.PartitionKey, body []byte, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error) {
			return azcosmos.ItemResponse{}, nil
		},
	}

	repo := NewWalletRepository(mock)

	err := repo.Create(context.Background(), wallet.Wallet{
		ID:        "1",
		AccountID: "acc1",
		Balance:   100,
	})

	assert.NoError(t, err)
}

func TestWalletRepository_GetByID(t *testing.T) {
	expected := wallet.Wallet{
		ID:        "1",
		AccountID: "acc1",
		Balance:   200,
	}

	data, _ := json.Marshal(expected)

	mock := &mockContainer{
		readFn: func(ctx context.Context, pk azcosmos.PartitionKey, id string, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error) {
			return azcosmos.ItemResponse{Value: data}, nil
		},
	}

	repo := NewWalletRepository(mock)

	result, err := repo.GetByID(context.Background(), "1", "acc1")

	assert.NoError(t, err)
	assert.Equal(t, expected, *result)
}
