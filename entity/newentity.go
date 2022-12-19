package entity

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "banker"
	password = "dhruv123"
	dbname   = "postgres"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func newBank(name string) {	

}

func newAccount(number int32, balance float32) {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// insert row
	x := `insert into "accounts" values($1,$2, $3, $4, $5)`
	_, e := db.Exec(x, 78, 5.3, 1234)
	CheckError(e)

	// close database
	defer db.Close()

}
