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
