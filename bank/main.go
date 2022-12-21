package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/balance", op.Menu())

	router.Run("localhost:8080")
}
