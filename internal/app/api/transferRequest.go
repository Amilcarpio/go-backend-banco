package api

type TransferRequest struct {
	Payer int     `json:"payer"`
	Payee int     `json:"payee"`
	Value float64 `json:"value"`
}
