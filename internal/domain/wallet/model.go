package wallet

type Wallet struct {
	ID        string `json:"id"`
	AccountID string `json:"accountId"`
	Balance   int64  `json:"balance"`
}
