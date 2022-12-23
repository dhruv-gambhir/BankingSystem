package setter_getter

import (
	"github.com/go-pg/pg"
	en "github.com/main/bank/entity"
)

func GetByID(dbRef *pg.DB, id int64) en.Account{
	var acc en.Account
	acc.ID = id
	acc.GetByID(dbRef)
	return acc
}

