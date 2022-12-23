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
}

func (x *Transaction) GetArrayBytID(db *pg.DB) {
	var trans []Transaction
	err := db.Model(&trans).Where("id = ?0", x.ID).Select()
	if err != nil {
		fmt.Println("Error getting transaction by account ID")
	}
}

func (x *Transaction) Update(db *pg.DB) {
	err := db.Update(x)
	if err != nil {
		fmt.Println("Error updating transaction")
	}
}

func (x *LoanTransaction) GetByID(db *pg.DB) {
	err := db.Select(x)
	if err != nil {
		fmt.Println("Error getting loan transaction by ID")
	}
}

func (x *Transaction) GetLArrayBytID(db *pg.DB) {
	var trans []LoanTransaction
	err := db.Model(&trans).Where("id = ?0", x.ID).Select()
	if err != nil {
		fmt.Println("Error getting transaction by account ID")
	}
}

func (x *LoanTransaction) Update(db *pg.DB) {
	err := db.Update(x)
	if err != nil {
		fmt.Println("Error updating loan transaction")
	}
	fmt.Println("Loan transaction updated for ID: ", x.LoanID)
}
