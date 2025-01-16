package models

type BalanceRequest struct {
	Address string `json:"address"`
}

type BalanceResponse struct {
	Balances map[string]float64 `json:"balances"`
}
