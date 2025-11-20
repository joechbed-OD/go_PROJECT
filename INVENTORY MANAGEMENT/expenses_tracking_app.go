package main

import (
	"fmt"
)

type expenses struct {
	ID      string
	Name    string
	Balance float64
	Cost    float64
}

type Spendings interface {
	// calculation()
	remainingBalance() float64
}

// func (e *expenses) calculation() error {
// 	calc := e.Balance - e.Cost
// 	if   e.Cost > e.remainingBalance(){
// 		return fmt.Errorf("Inssuficient Balance. Your current Balance is %.2f",e.Balance)
// 	}
// 	fmt.Println("")
// 	return nil
// }

func (e *expenses) remainingBalance() float64 {
	var total float64
	total = e.Balance - e.Cost
	fmt.Print("Remaining balance is ",total)
	return total
}

func main() {
	fmt.Println("EXPENSES TRACKING APP")
	fmt.Println("---------------------")
	expenditure := make(map[string]*expenses)

	for {
		fmt.Println("1.CHECK MY EXPENSES\n 2.CHECK MY REMAINING BALANCE")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var id, name string
			var balance, cost float64
			fmt.Println("CHECK MY EXPENSES")
			fmt.Println("ENTER ID FOR THE EXPENSES")
			fmt.Scan(&id)
			fmt.Println("ENTER WHAT THE EXPENSE IS")
			fmt.Scan(&name)
			fmt.Println("ENTER YOUR BALANCE")
			fmt.Scan(&balance)
			fmt.Println("ENTER COST OF EXPENSES")
			fmt.Scan(&cost)
			expen := &expenses{
				ID:      id,
				Name:    name,
				Balance: balance,
				Cost:    cost,
			}
			expenditure[id] = expen

			fmt.Println(expen.remainingBalance())
		case 2:
			var id string
			fmt.Println("ENTER ID TO CHECK BALANCE")
			fmt.Scan(&id)

			if expen, ok := expenditure[id]; ok {
				expen.remainingBalance()
			} else {
				fmt.Println("No expense found with that ID.")
			}
		default:
			fmt.Println("INVALID CHOICE")
		}
	}
}
