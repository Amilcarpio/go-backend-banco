package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {

	router := SetupRouter()
	router.Run(":8080")

}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	return r
}

func notImplementedHandler(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"error": fmt.Errorf("not implemented"),
	})
}
