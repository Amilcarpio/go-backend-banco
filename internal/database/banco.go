package database

import (
	"context"
	"fmt"

	"github.com/Amilcarpio/go-backend-banco/internal/app/models"
	"github.com/Amilcarpio/go-backend-banco/internal/app/services"
)

func CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (name, cpf, email, password, balance) VALUES (?, ?, ?, ?, ?)"

	_, err := db.ExecContext(ctx, query, user.FullName, user.CPF, user.Email, user.Password, user.Balance)

	if err != nil {
		return err
	}

	return nil
}

func CreateMerchant(ctx context.Context, merchant *models.Merchant) error {
	query := "INSERT INTO merchants (name, cnpj, email, password, balance) VALUES (?, ?, ?, ?, ?)"

	_, err := db.ExecContext(ctx, query, merchant.FullName, merchant.CNPJ, merchant.Email, merchant.Password, merchant.Balance)

	if err != nil {
		return err
	}

	return nil
}

func TransferMoney(ctx context.Context, senderID, receiverID int, amount float64) error {
	//iniciar a transação
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	//Obter os saldos atuais do remetente e destinatário
	var senderBalance, receiverBalance float64

	err = tx.QueryRowContext(ctx, "SELECT balance FROM users WHERE id = ?", senderID).Scan(&senderBalance)
	if err != nil {
		return err
	}

	err = tx.QueryRowContext(ctx, "SELECT balance FROM users WHERE id = ?", receiverID).Scan(&receiverBalance)
	if err != nil {
		return err
	}

	//verificar se o remetente tem saldo suficiente
	if senderBalance < amount {
		return fmt.Errorf("sender does not have enough balance")
	}

	//atualizar saldos
	newSenderBalance := senderBalance - amount
	newReceiverBalance := receiverBalance + amount

	//atualizar saldo do remetente
	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = ? WHERE id = ?", newSenderBalance, senderID)

	if err != nil {
		return err
	}

	//atualizar saldo do destinatário
	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = ? WHERE id = ?", newReceiverBalance, receiverID)

	if err != nil {
		return err
	}

	//commit da transação se tudo estiver correto
	err = tx.Commit()
	if err != nil {
		return err
	}

	//notificar o remetente e o destinatário
	err = services.SendNotification(ctx, senderID)
	if err != nil {
		return err
	}

	err = services.SendNotification(ctx, receiverID)
	if err != nil {
		return err
	}

	return nil
}
