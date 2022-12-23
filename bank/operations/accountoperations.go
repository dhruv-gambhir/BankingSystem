package operations

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

func NewTransaction() {

	fmt.Println("What would you like to do?")
	fmt.Println("1. Deposit money")
	fmt.Println("2. Withdraw money")
	fmt.Println("3. Transfer money")
	fmt.Println("4. Check balance")
	fmt.Println("5. Account Details")
	fmt.Println("6. Print Statement")
	fmt.Println("7. Exit")

}

func NewAccount(acc *en.Account) {

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

	//insert into database
	insertErr := db.Insert(&acc)
	if insertErr != nil {
		fmt.Println("Error inserting into table")
	} else {
		fmt.Println("Account inserted in Table")
	}

	closerr := db.Close()
	if closerr != nil {
		fmt.Println("Database connection failed")
		os.Exit(100)
	}
	fmt.Println("Database connection closed")

}
