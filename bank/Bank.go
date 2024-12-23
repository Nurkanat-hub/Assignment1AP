package bank

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (account *BankAccount) Deposit(amount float64) {
	account.Balance += amount
	fmt.Printf("Deposited: %.2f\n", amount)
}
func (account *BankAccount) Withdraw(amount float64) {
	if account.Balance >= amount {
		account.Balance -= amount
		fmt.Printf("Withdrew: %.2f\n", amount)
	} else {
		fmt.Println("Insufficient balance!")
	}
}
func (account *BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", account.Balance)
}
func Transaction(account *BankAccount, transactions []float64) {
	for _, transaction := range transactions {
		if transaction > 0 {
			account.Deposit(transaction)
		} else {
			account.Withdraw(-transaction)
		}
	}
}
func main() {
	account := BankAccount{
		AccountNumber: "123456789",
		HolderName:    "John Doe",
		Balance:       0.0,
	}

	var input string
	var amount float64

	for {
		fmt.Println("\nEnter an option:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get balance")
		fmt.Println("4. Exit")
		fmt.Print("Your choice: ")
		fmt.Scanln(&input)

		switch input {
		case "1":
			fmt.Print("Enter deposit amount: ")
			fmt.Scanln(&amount)
			account.Deposit(amount)
		case "2":
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scanln(&amount)
			account.Withdraw(amount)
		case "3":
			account.GetBalance()
		case "4":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
