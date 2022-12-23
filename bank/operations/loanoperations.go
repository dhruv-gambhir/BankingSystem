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

	/*
		var interest float64
		var amount float64
		var total_amount float64
		var installments int64
		var term int64

		loan.Installments = loan.Term
		interest = (amount * float64(term)) / 100
		total_amount = amount + interest
		loan.AmountPerInstallment = total_amount / float64(installments)
	*/

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
