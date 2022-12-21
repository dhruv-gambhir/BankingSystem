package operations

import (
	"fmt"
)

func Pay(loadID int){

}

func Check(loanID int){

}

func ManageLoan() {
	var choice int
	var loanID int
	fmt.Println("Welcome to the loan management menu")
	fmt.Println("Please enter your loan id")
	fmt.Scan(&loanID)

	for choice!=3 {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Pay Installment")
		fmt.Println("2. Check loan status")
		fmt.Println("3. Exit")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			Pay(loanID)
		case 2:
			Check(loanID)
		case 3:
			return
		default:
			fmt.Println("Invalid choice")
	}
	fmt.Println("Exiting loan management menu")
	}
}

func NewLoan() {

	

}