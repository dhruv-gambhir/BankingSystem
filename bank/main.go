package main

import (
	"github.com/gin-gonic/gin"
	op "github.com/main/bank/operations"

)

func main() {

	router := gin.Default()
	router.GET("/balance", op.Menu)
	
	router.Run("localhost:8080")
}
