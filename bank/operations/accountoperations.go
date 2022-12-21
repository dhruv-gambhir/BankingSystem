package operations

import (
	"fmt"
)

func Deposit(accountID int) {

}

func Withdraw(accountID int) {
}

func Transfer(accountID int) {

}

func Balance(accountID int) {
	
}

func Details(accountID int) {

}

func Statement(accountID int) {
	
}

func ManageAccount() {

	var choice int
	var accountID int
	var pin int
	fmt.Println("Welcome to the account management menu")
	fmt.Println("Please enter your account number")
	fmt.Scan(&accountID)
	fmt.Println("Please enter your pin")
	fmt.Scan(&pin)




	for choice!=7 {
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
			Deposit(accountID)
		case 2:
			Withdraw(accountID)
		case 3:
			Transfer(accountID)
		case 4:
			Balance(accountID)
		case 5:
			Details(accountID)
		case 6:
			Statement(accountID)
		case 7:
			return
		default:
			fmt.Println("Invalid choice")
		}
	fmt.Println("Exiting account management menu")
	}
}