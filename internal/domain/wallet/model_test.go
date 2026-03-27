package wallet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletCreation(t *testing.T) {
	wallet := Wallet{
		ID:        "wallet123",
		AccountID: "account456",
		Balance:   1000,
	}

	assert.Equal(t, "wallet123", wallet.ID)
	assert.Equal(t, "account456", wallet.AccountID)
	assert.Equal(t, int64(1000), wallet.Balance)
}
