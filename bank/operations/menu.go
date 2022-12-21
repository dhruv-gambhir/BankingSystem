package operations

import (
	"fmt"
	"github.com/main/bank/entity"
)

func Menu() {

	var choice int
	fmt.Println("Welcome to DiGi Bank")

	for (choice!=5) {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Manage Account")
		fmt.Println("2. Manage Loan")
		fmt.Println("3. New Account")
		fmt.Println("4. New Loan")
		fmt.Println("5. Exit")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			ManageAccount()
		case 2:
			ManageLoan()
		case 3:
			entity.NewAccount()
		case 4:
			entity.NewLoan()
		case 5:
			break
		default:
			fmt.Println("Invalid choice")
		}
	}
	fmt.Println("Thank you for using DiGi Bank")
}
