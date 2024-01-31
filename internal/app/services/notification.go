package services

import (
	"context"
	"fmt"
)

func SendNotification(ctx context.Context, userID int) error {
	fmt.Printf("Pagamento concluído. ID do usuário: %d\n", userID)

	return nil
}
