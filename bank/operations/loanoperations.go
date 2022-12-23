package operations

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/gin-gonic/gin"
	pg "github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

func Pay(loadID int) {

}

func Check(loanID int) {

}

func ManageLoan() {
	var choice int
	var loanID int
	fmt.Println("Welcome to the loan management menu")
	fmt.Println("Please enter your loan id")
	fmt.Scan(&loanID)

	for choice != 3 {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Pay Installment")
		fmt.Println("2. Check loan status")
		fmt.Println("3. Exit")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			Pay(loanID)
		case 2:
			Check(loanID)
		case 3:
			return
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("Exiting loan management menu")
	}
}

func NewLoan() {

	router := gin.Default()
	router.GET("/function", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the loan application menu \n Please enter the amount you would like to borrow \n Please enter the term of the loan in months",
		})
	})

	router.Run(":8080")

	var interest float64
	var amount float64
	var total_amount float64
	var installments int64
	var term int64

	fmt.Println("Welcome to the loan application menu")
	fmt.Println("Please enter the amount you would like to borrow")
	fmt.Scan(&amount)
	fmt.Println("Please enter the term of the loan in months")
	fmt.Scan(&term)
	installments = int64(term)
	interest = (amount * float64(term)) / 100
	total_amount = amount + interest
	ampi := total_amount / float64(installments)
	loanID := int64(rand.Int31())

	//push to database
	opts := &pg.Options{
		User:     "banker",
		Password: "dhruv123",
		Database: "postgres",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		fmt.Println("Database connection failed")
		os.Exit(100)
	}
	fmt.Println("Database connection successful")

	en.NewLoan(db, loanID, amount, term, ampi, installments)

	fmt.Println("Loan application successful")
	fmt.Println("Your loan id is", loanID)
	fmt.Println("Your loan amount is", amount, "rupees")
	fmt.Println("Your loan term is", term, "months")
	fmt.Println("Your loan interest is", interest, "rupees")
	fmt.Println("Your loan installments are", installments)
	fmt.Println("Your loan amount per installment is", ampi)

	closerr := db.Close()
	if closerr != nil {
		fmt.Println("Database connection failed")
		os.Exit(100)
	}
	fmt.Println("Database connection closed")

}
