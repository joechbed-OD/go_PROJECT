package main

import (
	"fmt"
	"time"
)

type Customer struct {
	ID string
	Name string
	Balance []float64
}

type checkType struct {
	withdraw float64
	deposit float64
}

type Transcation struct {
	Amount float64
	checkType
	Timestamp time.Time
}

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	GetBalance() float64
}



func main() {
	customers := make(map[string]*Customer)

	fmt.Println("Welcome to Go Bank CLI")

	// Ask for ID
	var id string
	fmt.Print("ENTER YOUR ID: ")
	fmt.Scan(&id)

	// Ask for Name
	var name string
	fmt.Print("ENTER YOUR NAME: ")
	fmt.Scan(&name)

	cust := &Customer{
		ID:      id,
		Name:    name,
		Balance: []float64{},
	}

	customers[id] = cust

	fmt.Printf("Customer created: ID=%s, Name=%s, Balance=%.2f\n", cust.ID, cust.Name, cust.GetBalance())

	fmt.Println("WELCOME TO THE GO BANK", cust.Name)
	fmt.Println("PLEASE SELECT AN OPTION BELOW:")

	fmt.Println("1. DEPOSIT")
	fmt.Println("2. WITHDRAW")
	fmt.Println("3. CHECK BALANCE")
	fmt.Println("4. EXIT")
	for {
		fmt.Println("\n--- MENU ---")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Exit")
		fmt.Print("Choose option: ")

		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("ENTER AMOUNT TO DEPOSIT: ")
			var amount float64
			fmt.Scan(&amount)
			cust.Deposit(amount)

		case 2:
			fmt.Print("ENTER AMOUNT TO WITHDRAW: ")
			var amount float64
			fmt.Scan(&amount)
			err := cust.Withdraw(amount)
			if err != nil {
				fmt.Println("Error:", err)
			}

		case 3:
			fmt.Printf("Current Balance: %.2f\n", cust.GetBalance())

		case 4:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option, try again.")
		}
	}

}

func (c *Customer) Deposit(amount float64) {
	c.Balance = append(c.Balance, amount)
	fmt.Printf("Deposited %.2f | New Balance: %.2f\n", amount, c.GetBalance())

}

func (c *Customer) Withdraw(amount float64) error {
	if amount > c.GetBalance() {
		return fmt.Errorf("insufficient funds: current balance is %.2f", c.GetBalance())
	}
	c.Balance = append(c.Balance, -amount)
	fmt.Printf("Withdrew %.2f | New Balance: %.2f\n", amount, c.GetBalance())
	return nil

}

func (c *Customer) GetBalance() float64 {
	total := 0.0
	for _, amount := range c.Balance {
		total += amount
	}
	return total
}