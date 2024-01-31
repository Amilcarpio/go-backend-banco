package services

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func CheckAuthorization(ctx context.Context) (bool, error) {
	//Url do serviço de autorização externo
	authorizationUrl := "https://run.mocky.io/v3/5794d450-d2e2-4412-8131-73d0293ac1cc"

	//Chamada HTTP GET para o serviço
	resp, err := http.Get(authorizationUrl)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	//Leitura da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("error reading response body: %v", err)
	}

	//Verifica se a resposta está autorizada
	if resp.StatusCode == http.StatusOK && string(body) == "true" {
		return true, nil
	}

	return false, nil
}
