package main

import (
	pg "github.com/main/bank/postgres"
	en "github.com/main/bank/entity"
)

func main() {

	pg.Connect()
	

	/*
		router := gin.Default()
		router.GET("/bank", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})

		router.Run(":8080")
	*/
}
