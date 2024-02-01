#!/bin/bash

# Configuração das variáveis de ambiente
export GIN_MODE=debug
export PORT=8000

# Iniciar a aplicação usando o modo de desenvolvimento do Gin
go run main.go