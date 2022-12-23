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
