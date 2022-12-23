package operations

import (
	"net/http"

	"github.com/gin-gonic/gin"

	en "github.com/main/bank/entity"
)

type LOGIN struct {
	Id  int64 `json:"id"`
	Pin int64 `json:"pin"`
}

var login LOGIN
var newAccount en.Account
var newLoan en.Loan
var newTransaction en.Transaction
var newLoanTransaction en.LoanTransaction

func Menu() {

	router := gin.Default()

	router.GET("/menu", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to DiGi Bank \n What would you like to do? \n 1. Transaction \n 2. Loan Transaction \n 3. New Account \n 4. New Loan \n 5. Display Account \n 6. Display Loan \n 7. Account Statement \n 8. Loan Statement",
		})
	})

	//new transaction

	router.GET("menu/function1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the new transaction menu ",
		})
	})

	router.POST("/menu/function1/login", func(c *gin.Context) {
		if err := c.BindJSON(&login); err != nil {
			return
		}
		if CheckLogin(&login) != 0 {
			c.JSON(200, gin.H{
				"message": "Login Successful",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Login Failed",
			})
			c.IndentedJSON(http.StatusCreated, login)
		}
	})

	router.POST("/menu/function1/set", func(c *gin.Context) {
		if err := c.BindJSON(&newTransaction); err != nil {
			return
		}
		NewTransaction(&newTransaction)
		c.IndentedJSON(http.StatusCreated, newTransaction)
	})

	//new loan transaction
	router.GET("menu/function2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the new loan transaction menu ",
		})
	})
	router.POST("/menu/function2/set", func(c *gin.Context) {
		if err := c.BindJSON(&newLoanTransaction); err != nil {
			return
		}

		NewLoanTransaction(&newLoanTransaction)
		c.IndentedJSON(http.StatusCreated, newLoanTransaction)
	})

	//new account
	router.GET("menu/function3", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the new account menu ",
		})
	})

	router.POST("/menu/function3/set", func(c *gin.Context) {
		if err := c.BindJSON(&newAccount); err != nil {
			return
		}
		//check errors
		if err := NewAccount(&newAccount); err != 0 {
			if err == 1 {
				c.JSON(200, gin.H{
					"message": "Cannot open database",
				})
			}
			if err == 2 {
				c.JSON(200, gin.H{
					"message": "Cannot insert into database",
				})
			}
			if err == 3 {
				c.JSON(200, gin.H{
					"message": "Cannot close database",
				})
			}
		}
		c.IndentedJSON(http.StatusCreated, newAccount)
	})

	//new loan
	router.GET("menu/function4", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the new loan menu ",
		})
	})
	router.POST("/menu/function4/set", func(c *gin.Context) {
		if err := c.BindJSON(&newLoan); err != nil {
			return
		}

		//set loan thinga
		var interest float64
		var total_amount float64
		var term int64
		var ampi float64

		newLoan.Installments = newLoan.Term
		interest = (newLoan.Amount * float64(term)) / 100
		total_amount = newLoan.Amount + interest
		ampi = total_amount / float64(newLoan.Installments)
		newLoan.AmountPerInstallment = ampi

		//check errors
		if err := NewLoan(&newLoan); err != 0 {
			if err == 1 {
				c.JSON(200, gin.H{
					"message": "Cannot open database",
				})
			}
			if err == 2 {
				c.JSON(200, gin.H{
					"message": "Cannot insert into database",
				})
			}
			if err == 3 {
				c.JSON(200, gin.H{
					"message": "Cannot close database",
				})
			}
		}
		c.IndentedJSON(http.StatusCreated, newLoan)
	})

	//display account
	router.POST("menu/function5", func(c *gin.Context) {
		if err := c.BindJSON(&login); err != nil {
			return
		}
		if CheckLogin(&login) != 0 {
			c.JSON(200, gin.H{
				"message": "Login Successful",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Login Failed",
			})
			c.IndentedJSON(http.StatusCreated, login)
		}

		c.JSON(200, GetAccount(login.Id))

	})

	//display loan
	router.POST("menu/function6", func(c *gin.Context) {
		if err := c.BindJSON(&login); err != nil {
			return
		}

		c.JSON(200, GetLoan(login.Id))
	})

	//display account transactions
	router.POST("menu/function7", func(c *gin.Context) {
		if err := c.BindJSON(&login); err != nil {
			return
		}
		if CheckLogin(&login) != 0 {
			c.JSON(200, gin.H{
				"message": "Login Successful",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Login Failed",
			})
			c.IndentedJSON(http.StatusCreated, login)
		}

		c.JSON(200, GetTransactions(login.Id))
	})

	//display loan transactions
	router.POST("menu/function8", func(c *gin.Context) {
		if err := c.BindJSON(&login); err != nil {
			return
		}

		c.JSON(200, GetLoanTransaction(login.Id))

	})

	//thanks
	router.GET("/menu/thanks", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Thank you for using DiGi Bank",
		})
	})

	// Listen and serve on
	router.Run(":8080")
}
