package main

import (
	op "github.com/main/bank/operations"
	pg "github.com/main/bank/postgres"
)

func main() {

	pg.Connect()
	op.Menu()

	/*
		router := gin.Default()
		router.GET("/bank", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})

		router.Run(":8080")
	*/
}
