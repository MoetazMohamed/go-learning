package main

import (
	"fmt"
)

func main1(){
	var revenue float64;
	var expenses float64;
	var taxrate float64;

	fmt.Println("Hello This Your Profit Calculator \n Please Enter the following")

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Println();

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Println();

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxrate)
	fmt.Println();
	
	earningBeforeTax :=  revenue - expenses

	fmt.Printf("Earnings BEFORE taxes: %.f \n" , earningBeforeTax)

	fmt.Printf("Earnings AFTER taxes %.f \n" , earningBeforeTax * (1- taxrate))

	fmt.Printf("Ratio: %0.2f \n",earningBeforeTax / (earningBeforeTax * (1- taxrate)))

	//    var i float64 = 0
	// 	for i < 10 {
	// 		fmt.Printf("value: %.f ", i);
	// 		i++;
	// 	}

}

func sum( x , y int) int{
     
	 return x + y 
}

func BeforeTax(revenue, expenses float64) float64{
	return revenue -  expenses

}

func AfterTax(revenue, expenses, taxes float64) (before float64 , place int) {
	before =  BeforeTax(revenue, expenses) *(1- taxes) 
	place = 1
	return 
}
 
func Ratio(revenue, expenses, taxes float64) float64{
	x , _ := AfterTax(revenue, expenses, taxes)
	return BeforeTax(revenue, expenses) / x
}