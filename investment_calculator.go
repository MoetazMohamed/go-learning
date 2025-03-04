package main

// this means that this class belongs to a package called main
// No file can exist without being part of a package same idea as namespace or package in java
// why main as name of package
// go run app.go ==> you are running app
// go build ==> create executable file that can be executed on any system
//  main module ==> neds to have oackage main to know which packaghe is the entry point

// it tells go that main will be the main entry package when building our app
import (
	"fmt"
	"math"
) // "fmt" is a package that we import to use

// very similar to the public static void main(String[] args)
// only functions that start with capital letter is similar to export 
func main2() {
	fmt.Print("Hello World")
	// most code must be in functions
	// one main function should exist in a package because if we have two mains whoch one will execute when we run go ???
	// go PACKAGE does not necessary needs main function.. some times u dont need to run it but u just need create function that will be used in
	// other files
	const inflationRate float64 = 2
	investmentAmount, years := 1000.0 , 10.0 // infer the type of the value 
	 expectedReturnRate  := 5.5 // --> we dont need explicitly to put the type we can commit var and then use :=
	// fmt.Scan(&investmentAmount, &expectedReturnRate, &years)

	fmt.Print("Please input investment amount ")
	fmt.Scan(&investmentAmount)
	fmt.Println()
	fmt.Print("Please input years ")
	fmt.Scan(&investmentAmount)
	fmt.Println()
	fmt.Print("Please input expeted returns ")
	
	fmt.Scan(&investmentAmount)
	fmt.Println()
	// if u want to make the type inferred then u dont need to have var anymore 
	// go is statically typed language
	var futureValue = math.Pow(investmentAmount*(expectedReturnRate/100+1), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print(futureValue)
	fmt.Print(futureRealValue)
}
