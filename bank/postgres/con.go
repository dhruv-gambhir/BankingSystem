package postgres

import (
	"fmt"
	"os"

	pg "github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

func Connect() {

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

	en.NewAccountTable(db)
	en.NewLoanTable(db)
	en.NewTransactionTable(db)
	en.NewLoanTransactionTable(db)

	closerr := db.Close()
	if closerr != nil {
		fmt.Println("Database connection failed")
		os.Exit(100)
	}
	fmt.Println("Database connection closed")

}
