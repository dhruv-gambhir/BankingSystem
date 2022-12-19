package setter_getter

import (
	"database/sql"
	"fmt"
)

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func getBalance(a int) float64 {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// get row
	x := `select * from accounts where account_number = $1`
	_, e := db.Exec(x, a)
	CheckError(e)

	// close database
	defer db.Close()
	var balance float64

	return balance
}
