package entity

type Account struct {
	tablename struct{} `sql:"accounts"`
	Bank      string   `json:"bank" sql:"bank"`
	Branch    string   `json:"branch" sql:"branch"`
	ID        int64    `json:"id" sql:"id"`
	Name      string   `json:"name" sql:"name"`
	Pin       int64    `json:"pin" sql:"pin"`
	Balance   float64  `json:"balance" sql:"balance"`
}

type Loan struct {
	tablename            struct{} `sql:"loans"`
	LoanID               int64    `json:"id" sql:"id"`
	Amount               float64  `json:"amount" sql:"amount"`
	Term                 int64    `json:"term" sql:"term_in_months"`
	AmountPerInstallment float64  `json:"amount_per_installment" sql:"amount_per_installment"`
	Installments         int64    `json:"installments" sql:"installments"`
	AmountPayed          float64  `json:"payed" sql:"payed"`
	InstallmentsPayed    int64    `json:"installments_payed" sql:"installments_payed"`
}

type Transaction struct {
	tablename struct{} `sql:"transactions"`
	ID        int64    `json:"id" sql:"id"`
	Amount    float64  `json:"amount" sql:"amount"`
	Type      string   `json:"type" sql:"type"`
}

type LoanTransaction struct {
	tablename    struct{} `sql:"loan_transactions"`
	Amount       float64  `json:"amount" sql:"amount"`
	LoanID       int64    `json:"loan_id" sql:"loan_id"`
	Installments int64    `json:"installment" sql:"installment"`
}
