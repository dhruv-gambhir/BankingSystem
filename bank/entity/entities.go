package entity

type account struct {
	tablename struct{} `db:"accounts"`
	Bank      string   `json:"bank" db:"bank"`
	Branch    string   `json:"branch" db:"branch"`
	Id        int64    `json:"id" db:"id"`
	Name      string   `json:"name" db:"name"`
	Pin       string   `json:"pin" db:"pin"`
}

type loan struct {
	tablename         struct{} `db:"loans"`
	LoanID            int64    `json:"id" db:"id"`
	Amount            int64    `json:"amount" db:"amount"`
	Term              int64    `json:"term" db:"term"`
	Installments      int64    `json:"installments" db:"installments"`
	AmountPayed       int64    `json:"payed" db:"payed"`
	InstallmentsPayed int64    `json:"installments_payed" db:"installments_payed"`
}

type transaction struct {
	tablename struct{} `db:"transactions"`
	LoanID    int64    `json:"loan_id" db:"loan_id"`
	Amount    int64    `json:"amount" db:"amount"`
}

type loan_transaction struct {
	tablename struct{} `db:"loan_transactions"`
	ID        int64    `json:"id" db:"id"`
	Amount    int64    `json:"amount" db:"amount"`
	LoanID    int64    `json:"loan_id" db:"loan_id"`
}
