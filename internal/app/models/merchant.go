package models

type Merchant struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	CNPJ     string `json:"cnpj"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
}
