package api

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//configurar rotas de api
	SetupRoutes(r)

	r.Run(":8080")
}
