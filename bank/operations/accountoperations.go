package operations

import (
	"fmt"
)

func ManageAccount() {

	var choice int
	fmt.Println("Welcome to the account management menu")
	fmt.Println("Please enter your account number")


	for ;; {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Deposit money")
		fmt.Println("2. Withdraw money")
		fmt.Println("3. Transfer money")
		fmt.Println("4. Check balance")
		fmt.Println("5. Account Details")
		fmt.Println("6. Print Statement")
		fmt.Println("7. Exit")

		fmt.Scanln(&choice)
		switch choice {
		case 1:
			Deposit()
		case 2:
			Withdraw()
		case 3:
			Transfer()
		case 4:
			Balance()
		case 5:
			Details()
		case 6:
			Statement()
		case 7:
			return
		default:
			fmt.Println("Invalid choice")
	}
	fmt.Println("Exiting account management menu")
}