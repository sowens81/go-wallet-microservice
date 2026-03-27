package wallet

import "context"

type Repository interface {
	Create(ctx context.Context, wallet Wallet) error
	GetByID(ctx context.Context, id string, accountID string) (*Wallet, error)
}
