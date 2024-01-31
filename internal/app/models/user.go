package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
}
