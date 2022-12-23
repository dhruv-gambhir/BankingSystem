package main

import (
	op "github.com/main/bank/operations"
	pg "github.com/main/bank/postgres"
)

func main() {

	pg.Connect()
	op.Menu()
}
