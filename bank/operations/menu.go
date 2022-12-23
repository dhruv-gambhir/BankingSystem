package operations

import (
	"net/http"

	"github.com/gin-gonic/gin"

	en "github.com/main/bank/entity"
)

/*
type INPUT struct {
	Choice string `json:"choice"`
}

var input INPUT
var choice int
*/

var newAccount en.Account
var newLoan en.Loan
var newTransaction en.Transaction
var newLoanTransaction en.LoanTransaction

func Menu() {

	router := gin.Default()

	router.GET("/menu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to DiGi Bank \n What would you like to do? \n 1. Manage Account \n 2. Manage Loan \n 3. New Account \n 4. New Loan",
		})
	})
	/*
		router.POST("/menu/choice", func(c *gin.Context) {
			if err := c.BindJSON(&input); err != nil {
				return
			}
			choice, _ = strconv.Atoi(input.Choice)
			c.IndentedJSON(http.StatusCreated, input)
		})

		router.GET("/menu/getchoice", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": choice,
			})
		})
	*/

	//manage account
	router.GET("menu/function1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the manage account menu ",
		})
	})

	router.POST("/menu/function1/set", func(c *gin.Context) {
		if err := c.BindJSON(&newTransaction); err != nil {
			return
		}
		NewTransaction(&newTransaction)
		c.IndentedJSON(http.StatusCreated, newAccount)
	})
	ManageAccount()

	//manage loan
	router.GET("menu/function2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the manage loan menu ",
		})
	})
	ManageLoan()

	//new account
	router.GET("menu/function3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the new account menu ",
		})
	})

	router.POST("/menu/function/set", func(c *gin.Context) {
		if err := c.BindJSON(&newAccount); err != nil {
			return
		}
		NewAccount(&newAccount)
		c.IndentedJSON(http.StatusCreated, newAccount)
	})

	//new loan
	router.GET("menu/function4", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the new loan menu ",
		})
	})
	NewLoan()

	//thanks
	router.GET("/menu/thanks", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Thank you for using DiGi Bank",
		})
	})

	// Listen and serve on
	router.Run(":8080")
}
