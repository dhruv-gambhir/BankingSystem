package entity

import (
	"fmt"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

func NewAccount(db *pg.DB) {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&account{}, opts)
	if createErr != nil {
		fmt.Println("Error creating table")
	}
}

func NewLoan() {

}

func NewTransaction() {

}

func NewLoanTransaction() {

}
