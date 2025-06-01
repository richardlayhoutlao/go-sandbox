package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64


	fmt.Print("Enter Revenue:\n")
	fmt.Scan(&revenue)
	fmt.Print("Enter Expenses:\n")
	fmt.Scan(&expenses)
	fmt.Print("Enter Tax Rate:\n")
	fmt.Scan(&taxRate)

	//EBT
	EBT := revenue - expenses
	fmt.Print("EBT (Earnings Before Tax): ", EBT, "\n")

	//EAT
	profit := EBT * (1 - taxRate/ 100)
	fmt.Print("EAT (Earnings After Tax): ", profit, "\n")

	//Ratio
	ratio := EBT / profit
	fmt.Print("Profit: ", ratio, "\n")

}
