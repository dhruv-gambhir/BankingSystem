package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	fmt.Println("Enter account number: ")
	router := gin.Default()
	router.GET("/balance", getBalance())

	router.Run("localhost:8080")
}
