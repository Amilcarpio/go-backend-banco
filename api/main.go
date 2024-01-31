package main

import (
	"github.com/Amilcarpio/go-backend-banco/internal/app"
)

func main() {
	router := app.SetupRouter()
	router.Run(":8080")
}
