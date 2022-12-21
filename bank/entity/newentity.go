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
	createErr := db.CreateTable(&account{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}
}

func NewLoanTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&loan{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}

}

func NewTransactionTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&transaction{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}
}

func NewLoanTransactionTable(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&loan_transaction{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}

}

func NewAccount(db *pg.DB, accountID int64, pin int, name string, balance float64, bank string, branch string) {
	var acc account
	acc.ID = accountID
	acc.Pin = pin
	acc.Balance = balance
	acc.Name = name
	acc.Bank = bank
	acc.Branch = branch
	insertErr := db.Insert(&acc)
	if insertErr != nil {
		fmt.Println("Error inserting into table")
	}
}
