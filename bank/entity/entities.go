package entity

type account struct {
	tablename struct{} `sql:"accounts"`
	Bank      string   `json:"bank" sql:"bank"`
	Branch    string   `json:"branch" sql:"branch"`
	ID        int64    `json:"id" sql:"id"`
	Name      string   `json:"name" sql:"name"`
	Pin       int      `json:"pin" sql:"pin"`
	Balance   float64  `json:"balance" sql:"balance"`
}

type loan struct {
	tablename         struct{} `sql:"loans"`
	LoanID            int64    `json:"id" sql:"id"`
	Amount            int64    `json:"amount" sql:"amount"`
	Term              int64    `json:"term" sql:"term"`
	Installments      int64    `json:"installments" sql:"installments"`
	AmountPayed       int64    `json:"payed" sql:"payed"`
	InstallmentsPayed int64    `json:"installments_payed" sql:"installments_payed"`
}

type transaction struct {
	tablename struct{} `sql:"transactions"`
	ID        int64    `json:"id" sql:"id"`
	Amount    int64    `json:"amount" sql:"amount"`
}

type loan_transaction struct {
	tablename struct{} `sql:"loan_transactions"`
	ID        int64    `json:"id" sql:"id"`
	Amount    int64    `json:"amount" sql:"amount"`
	LoanID    int64    `json:"loan_id" sql:"loan_id"`
}
