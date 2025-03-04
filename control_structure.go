package main

import (
	"fmt"
	"go-learning/fileops"
)

const balanceFile = "balance.txt"

func main(){
	balance , err := fileops.ReadBalanceFromFile(balanceFile)
	if err != nil {
		fmt.Println("ERROR: %v", err)
		fmt.Println(err)
		//panic("Can't continue, sorry.")
	}

	for {
		var yourChoice int ;
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance\n2. Deposit money\n3. Withdraw money\n4. Exit")
		fmt.Print("Your choice: ")
		fmt.Scan(&yourChoice)
		switch yourChoice {
		case 1 :
			fmt.Printf("Your balance is %v\n", balance)
		case 2: 
			fmt.Print("deposit amount: ")
			var depositAmount float64;
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount")
				continue
			}
			balance += depositAmount
			fileops.WriteBalanceToFile(float64(balance), balanceFile)
			
		case 3:
			fmt.Print("withdrawal amount: ")
			var withdrawAmount float64;
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount")
				continue
			}
			balance -= withdrawAmount
			fileops.WriteBalanceToFile(float64(balance), balanceFile)
			
		case 4 :
			fmt.Println("Thank you bye bye")
			return
		default :
		fmt.Println("Invalid input please retry")
		continue
		}
	}
}