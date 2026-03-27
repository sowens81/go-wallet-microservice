package config

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/sowens81/go-wallet-microservice/internal/shared/ids"
)

func NewCosmosContainerWithAzureCredential(
	ctx context.Context,
	endpoint, dbName, containerName string,
) (*azcosmos.ContainerClient, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	client, err := azcosmos.NewClient(endpoint, cred, nil)
	if err != nil {
		return nil, err
	}

	db, dbCreated, err := ensureDatabase(ctx, client, dbName)
	if err != nil {
		return nil, err
	}

	container, containerCreated, err := ensureContainer(ctx, db, containerName)
	if err != nil {
		return nil, err
	}

	// 👉 Seed ONLY if fresh
	if dbCreated || containerCreated {
		if err := seedWallets(ctx, container); err != nil {
			return nil, err
		}
	}

	return container, nil
}

func ensureDatabase(ctx context.Context, client *azcosmos.Client, dbName string) (*azcosmos.DatabaseClient, bool, error) {
	props := azcosmos.DatabaseProperties{ID: dbName}

	_, err := client.CreateDatabase(ctx, props, nil)
	if err != nil {
		if isConflictError(err) {
			db, err := client.NewDatabase(dbName)
			return db, false, err
		}
		return nil, false, err
	}

	db, err := client.NewDatabase(dbName)
	return db, true, err
}

func ensureContainer(ctx context.Context, db *azcosmos.DatabaseClient, containerName string) (*azcosmos.ContainerClient, bool, error) {
	props := azcosmos.ContainerProperties{
		ID: containerName,
		PartitionKeyDefinition: azcosmos.PartitionKeyDefinition{
			Paths: []string{"/accountId"},
			Kind:  azcosmos.PartitionKeyKindHash,
		},
	}

	_, err := db.CreateContainer(ctx, props, nil)
	if err != nil {
		if isConflictError(err) {
			c, err := db.NewContainer(containerName)
			return c, false, err
		}
		return nil, false, err
	}

	c, err := db.NewContainer(containerName)
	return c, true, err
}

func isConflictError(err error) bool {
	var respErr *azcore.ResponseError
	if errors.As(err, &respErr) {
		return respErr.StatusCode == 409
	}
	return false
}

func seedWallets(ctx context.Context, container *azcosmos.ContainerClient) error {
	for i := 1; i <= 10; i++ {

		id := ids.GenerateShortID()
		accountID := ids.GenerateShortID()

		wallet := map[string]interface{}{
			"id":        id,
			"accountId": accountID,
			"balance":   int64(100 * i),
		}

		data, err := json.Marshal(wallet)
		if err != nil {
			return err
		}

		pk := azcosmos.NewPartitionKeyString(wallet["accountId"].(string))

		_, err = container.CreateItem(ctx, pk, data, nil)
		if err != nil {
			// Ignore duplicates (if rerun)
			if !isConflictError(err) {
				return err
			}
		}
	}

	return nil
}
