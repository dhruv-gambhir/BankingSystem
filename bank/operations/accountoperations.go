package operations

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

func Deposit(accountID int) {

}

func Withdraw(accountID int) {
}

func Transfer(accountID int) {

}

func Balance(accountID int) {

}

func Details(accountID int) {

}

func Statement(accountID int) {

}

func ManageAccount() {

	var choice int
	var accountID int
	var pin int
	fmt.Println("Welcome to the account management menu")
	fmt.Println("Please enter your account number")
	fmt.Scan(&accountID)
	fmt.Println("Please enter your pin")
	fmt.Scan(&pin)

	for choice != 7 {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Deposit money")
		fmt.Println("2. Withdraw money")
		fmt.Println("3. Transfer money")
		fmt.Println("4. Check balance")
		fmt.Println("5. Account Details")
		fmt.Println("6. Print Statement")
		fmt.Println("7. Exit")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			Deposit(accountID)
		case 2:
			Withdraw(accountID)
		case 3:
			Transfer(accountID)
		case 4:
			Balance(accountID)
		case 5:
			Details(accountID)
		case 6:
			Statement(accountID)
		case 7:
			return
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("Exiting account management menu")
	}
}

func NewAccount() {
	var choice int
	var bankName string
	var branchName string
	var id int64
	var name string
	var pin int
	var balance float64

	//select bank
	fmt.Println("Welcome to the new account menu")
	fmt.Println("Please select a bank")
	fmt.Println("1. ICICI")
	fmt.Println("2. HDFC")
	fmt.Println("3. SBI")
	fmt.Println("4. BOB")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		bankName = "ICICI"
	case 2:
		bankName = "HDFC"
	case 3:
		bankName = "SBI"

	case 4:
		bankName = "BOB"
	default:
		fmt.Println("Invalid choice")
	}

	//select branch
	fmt.Println("Please select a branch")
	fmt.Println("1. Mumbai")
	fmt.Println("2. Delhi")
	fmt.Println("3. Bangalore")
	fmt.Println("4. Chennai")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		branchName = "Mumbai"
	case 2:
		branchName = "Delhi"
	case 3:
		branchName = "Bangalore"
	case 4:
		branchName = "Chennai"
	default:
		fmt.Println("Invalid choice")
	}

	//enter name
	fmt.Println("Please enter your name")
	fmt.Scan(&name)

	//enter pin
	fmt.Println("Please enter your pin")
	fmt.Scan(&pin)

	//enter balance
	fmt.Println("Please enter the amount you want to deposit")
	fmt.Scan(&balance)

	//create account
	fmt.Println("Creating account...")
	fmt.Println("Account created successfully")
	id = int64(rand.Int31())
	fmt.Println("Your account number is", id)
	fmt.Println("Your account balance is", balance)
	fmt.Println("Your account pin is", pin)
	fmt.Println("Your account name is", name)
	fmt.Println("Your account bank is", bankName)
	fmt.Println("Your account branch is", branchName)

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

	en.NewAccount(db, id, pin, name, balance, bankName, branchName)

	closerr := db.Close()
	if closerr != nil {
		fmt.Println("Database connection failed")
		os.Exit(100)
	}
	fmt.Println("Database connection closed")

}
