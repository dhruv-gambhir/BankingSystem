package operations

import (
	pg "github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

var ltcount int64 = 1

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

	ln := &en.Loan{LoanID: ltran.LoanID}
	ln.GetByID(db)
	ln.InstallmentsPayed += 1
	ln.AmountPayed += ltran.Amount
	ln.Update(db)

	ltran.TransactionID = ltcount
	ltran.Amount = ln.AmountPerInstallment
	ltcount++
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

func GetLoanTransaction(id int64) interface{} {
	opts := &pg.Options{
		User:     "banker",
		Password: "dhruv123",
		Database: "postgres",
		Addr:     "localhost:5432",
	}

	var db *pg.DB = pg.Connect(opts)

	var x []en.LoanTransaction
	db.Model(&x).Where("loan_id = ?", id).Select()

	db.Close()

	return x
}
