package entity

import (
	"fmt"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

func NewAccountTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Account{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}
}

func NewLoanTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Loan{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}

}

func NewTransactionTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&Transaction{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}
}

func NewLoanTransactionTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&LoanTransaction{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}

}
