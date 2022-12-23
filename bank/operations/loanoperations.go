package operations

import (
	pg "github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

func NewLoanTransaction(ltran *en.LoanTransaction) int {
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

	//insert into database
	insertErr := db.Insert(ltran)
	if insertErr != nil {
		return 2
	}

	closerr := db.Close()
	if closerr != nil {
		return 3
	}

	return 0
}

func NewLoan(loan *en.Loan) int {

	//connect to database
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

	//insert to database
	insertErr := db.Insert(loan)
	if insertErr != nil {
		return 2
	}

	closerr := db.Close()
	if closerr != nil {
		return 3
	}

	return 0

}

func GetLoan(id int64) en.Loan {
	opts := &pg.Options{
		User:     "banker",
		Password: "dhruv123",
		Database: "postgres",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)

	ln := &en.Loan{LoanID: login.Id}
	ln.GetByID(db)

	db.Close()

	return *ln

}

func GetLoanTransaction(id int64) []en.LoanTransaction {
	opts := &pg.Options{
		User:     "banker",
		Password: "dhruv123",
		Database: "postgres",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)

	ltran := &en.LoanTransaction{LoanID: id}

	ltran.GetByID(db)

	db.Close()

	return []en.LoanTransaction{*ltran}
}
