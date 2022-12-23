package operations

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

func NewTransaction(tran *en.Transaction) {

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

	if tran.Type == "credit" {
		acc := &en.Account{ID: tran.ID}
		acc.GetByID(db)
		acc.Balance += tran.Amount
		acc.Update(db)
	} else if tran.Type == "debit" {
		acc := &en.Account{ID: tran.ID}
		acc.GetByID(db)
		acc.Balance -= tran.Amount
		acc.Update(db)
	}

	//insert into database
	insertErr := db.Insert(tran)
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

func NewAccount(acc *en.Account) int {

	//push to database
	opts := &pg.Options{
		User:     "banker",
		Password: "dhruv123",
		Database: "postgres",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		return 1
	}
	fmt.Println("Database connection successful")

	//insert into database
	insertErr := db.Insert(acc)
	if insertErr != nil {
		return 2
	}

	closerr := db.Close()
	if closerr != nil {
		return 3
	}

	return 0
}

func CheckLogin(log *LOGIN) int {

	//var acc en.Account

	opts := &pg.Options{
		User:     "banker",
		Password: "dhruv123",
		Database: "postgres",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)

	//select from database
	//db.Model(acc).Where("id = ?", login.Id).Select()

	acc := &en.Account{ID: login.Id}
	acc.GetByID(db)

	if acc.Pin == login.Pin {
		return 1
	}

	db.Close()

	return 0

}

func GetAccount(id int64) en.Account {
	opts := &pg.Options{
		User:     "banker",
		Password: "dhruv123",
		Database: "postgres",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)

	acc := &en.Account{ID: login.Id}
	acc.GetByID(db)

	db.Close()

	return *acc

}
