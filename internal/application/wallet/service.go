package wallet

import (
	"context"

	domain "github.com/sowens81/go-wallet-microservice/internal/domain/wallet"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateWallet(ctx context.Context, wallet domain.Wallet) error {
	// business logic could go here
	return s.repo.Create(ctx, wallet)
}

func (s *Service) GetWallet(ctx context.Context, id, accountID string) (*domain.Wallet, error) {
	return s.repo.GetByID(ctx, id, accountID)
}
