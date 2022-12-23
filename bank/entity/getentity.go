package entity

import (
	"fmt"

	"github.com/go-pg/pg"
)

func (x *Account) GetByID(db *pg.DB) {
	err := db.Select(x)
	if err != nil {
		fmt.Println("Error getting account by ID")
	}
	fmt.Println("Account found for ID: ", x.ID)
}

func (x *Account) Update(db *pg.DB) {
	err := db.Update(x)
	if err != nil {
		fmt.Println("Error updating account")
	}
	fmt.Println("Account updated for ID: ", x.ID)
}

func (x *Loan) GetByID(db *pg.DB) {
	err := db.Select(x)
	if err != nil {
		fmt.Println("Error getting loan by ID")
	}
	fmt.Println("Loan found for ID: ", x.LoanID)
}

func (x *Loan) Update(db *pg.DB) {
	err := db.Update(x)
	if err != nil {
		fmt.Println("Error updating loan")
	}
	fmt.Println("Loan updated for ID: ", x.LoanID)
}

func (x *Transaction) GetByID(db *pg.DB) {
	err := db.Select(x)
	if err != nil {
		fmt.Println("Error getting transaction by ID")
	}
	fmt.Println("Transaction found for ID: ", x.ID)
}

func (x *Transaction) Update(db *pg.DB) {
	err := db.Update(x)
	if err != nil {
		fmt.Println("Error updating transaction")
	}
	fmt.Println("Transaction updated for ID: ", x.ID)
}

func (x *LoanTransaction) GetByID(db *pg.DB) {
	err := db.Select(x)
	if err != nil {
		fmt.Println("Error getting loan transaction by ID")
	}
	fmt.Println("Loan transaction found for ID: ", x.LoanID)
}

func (x *LoanTransaction) Update(db *pg.DB) {
	err := db.Update(x)
	if err != nil {
		fmt.Println("Error updating loan transaction")
	}
	fmt.Println("Loan transaction updated for ID: ", x.LoanID)
}
