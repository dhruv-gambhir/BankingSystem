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

func NewAccount(db *pg.DB, accountID int64, pin int64, name string, balance float64, bank string, branch string) {
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
	} else {
		fmt.Println("Account inserted in Table")
	}
}

func NewLoan(db *pg.DB, loanID int64, amount float64, term int64, ampi float64, installments int64) {
	var loan loan
	loan.LoanID = loanID
	loan.Amount = amount
	loan.Term = term
	loan.AmountPerInstallment = ampi
	loan.Installments = installments
	loan.AmountPayed = float64(0)
	loan.InstallmentsPayed = int64(0)

	insertErr := db.Insert(&loan)
	if insertErr != nil {
		fmt.Println("Error inserting into table")
	} else {
		fmt.Println("Loan inserted in Table")
	}
}

func NewTransaction(db *pg.DB, transactionID int64, amount float64, accountID int64) {
	var transaction transaction
	transaction.ID = transactionID
	transaction.Amount = amount
	transaction.ID = accountID
	insertErr := db.Insert(&transaction)
	if insertErr != nil {
		fmt.Println("Error inserting into table")
	} else {
		fmt.Println("Transaction inserted in Table")
	}
}

func NewLoanTransaction(db *pg.DB, installments int64, amount float64, loanID int64) {
	var transaction loan_transaction
	transaction.Installments = installments
	transaction.Amount = amount
	transaction.LoanID = loanID
	insertErr := db.Insert(&transaction)
	if insertErr != nil {
		fmt.Println("Error inserting into table", insertErr)
	} else {
		fmt.Println("Loan Transaction inserted in Table")
	}
}
