package cosmos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sowens81/go-wallet-microservice/internal/domain/wallet"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

type ContainerAPI interface {
	CreateItem(ctx context.Context, pk azcosmos.PartitionKey, body []byte, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error)
	ReadItem(ctx context.Context, pk azcosmos.PartitionKey, id string, o *azcosmos.ItemOptions) (azcosmos.ItemResponse, error)
}

type walletRepository struct {
	container ContainerAPI
}

func NewWalletRepository(container ContainerAPI) wallet.Repository {
	return &walletRepository{container: container}
}

func (r *walletRepository) Create(ctx context.Context, w wallet.Wallet) error {
	data, err := json.Marshal(w)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}

	pk := azcosmos.NewPartitionKeyString(w.AccountID)

	_, err = r.container.CreateItem(ctx, pk, data, nil)
	if err != nil {
		return fmt.Errorf("create item: %w", err)
	}

	return nil
}

func (r *walletRepository) GetByID(ctx context.Context, id, accountID string) (*wallet.Wallet, error) {
	pk := azcosmos.NewPartitionKeyString(accountID)

	resp, err := r.container.ReadItem(ctx, pk, id, nil)
	if err != nil {
		return nil, fmt.Errorf("read item: %w", err)
	}

	var w wallet.Wallet
	if err := json.Unmarshal(resp.Value, &w); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return &w, nil
}
