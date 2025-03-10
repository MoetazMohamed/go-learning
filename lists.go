package main

import (
	"fmt"
	"errors"
	
)

type Product struct {
	title string
	id    string
	price float64
}

func main6() {
	// how to create arrays
	// 1. prdarr := [length]type {value1,  value2}
	// 2. var productNames []string
	// var prdarr []int
	// for i:= 0 ; i<10 ; i++{
		
	// 	prdarr = append(prdarr, i)
		
	// }

	// fmt.Println(prdarr)

	// // binary search in go :)
	// index , err := binarySearch(prdarr, 10)
	// if err != nil{
	// 	fmt.Println(err)
		
	// } else {
	// 	fmt.Printf("target is found at index %v", index)
	// }


	// //var arr []int = []int {1,2,3,4}

	
	// fixedSizeArr := [5]int{}
	// for i := 0 ; i <len(fixedSizeArr); i++{
	// 	fixedSizeArr[i] = i*i
	// }
	// fmt.Println(fixedSizeArr)
	// slicedArr := fixedSizeArr[3:]
	// for i := 0 ;i < len(slicedArr) ; i++{
	// 	fmt.Printf("value: %v \n", slicedArr[i])
	// }

	// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
	hobbies := [3]string{"soccer", "coding", "hockey"}
	
	fmt.Println(hobbies)
	
	
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
	fmt.Printf("first element of the list : %v\n",hobbies[0])
	slicedArr :=  hobbies[0:2]
	fmt.Printf("first and second elements in new list : %v\n",slicedArr)
// 3) Create a slice based on the first element that contains
//		the first and second elements.,
//		Create that slice in two different ways (i.e. create two slices in the end)
	slicedArr = hobbies[:1]
	fmt.Printf("first way %v\n",slicedArr)
	var slicedArr2 []string
	slicedArr2 = append(slicedArr2, (hobbies[:1])...)
	fmt.Printf("second way %v\n",slicedArr2)

// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
fmt.Println(slicedArr[0:3])

// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"pass interview", "do an application with go"}
	fmt.Printf("Initialized course goals %v\n", courseGoals)
	courseGoals[1] =  "learn go routines"
	fmt.Printf("changed secoind course goals %v\n", courseGoals)
	courseGoals = append(courseGoals, "learn lists in go")
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
courseGoals[1] =  "learn go routines"
fmt.Printf("Changed secoind course goals %v\n", courseGoals)
courseGoals = append(courseGoals, "learn lists in go")
fmt.Printf("Add new element to the slice %v\n", courseGoals)

// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
	products := []Product {
		Product{
			title: "energy",
			id: "1",
			price : 100,
		}, 
		Product{
			title: "orange",
			id: "2",
			price : 10,
		},
	}
	products = append(products, Product{
		title: "shs",
		id : "3",
		price: 23,
	})
	fmt.Println(products)

	index1, index2, err := TwoSum([]int{1,2,3,4}, 10)

	if(err != nil){
		fmt.Println(err)
	}else {
		fmt.Println(index1, index2)
	}

     strMap := map[string]string{}
	 strMap["google"] = "gg"
	 strMap["facebook"] = "ff"
	 strMap["instgram"] = "ii"
	 for key , val :=  range strMap{
		fmt.Println(key, val)
	 } 

	 for i , val := range slicedArr {
		fmt.Println(i , val)
	 }
}
// 2 3 4 5 6 7 8    7 - 2 / 2 = 5 /2 = 2+ 2
func binarySearch(arr []int , target int) (int , error){
	left:= 0
	right:= len(arr) -1 
	for left < right {
		// get the mid point 
		// if midpoint greater than target then move left to midpoint 
		// if midpoint less than the target then move right to midpoint
		midpoint := left + (right - left ) / 2 

		if(target == arr[midpoint]){
			return midpoint , nil
		} 

		if(target > arr[midpoint]){
			left = midpoint + 1
		} else {
			right = midpoint -1
		}
	}

	return -1, errors.New("number is not found")
}

func TwoSum(arr []int, target int) (int , int , error){
	dict :=  map[int]int{}

	for i:= 0 ; i < len(arr) ; i++{
		val, exist := dict[target - arr[i]] 
		if !exist{
			dict[arr[i]] = i
		} else {
			return val, i , nil
		}
		
		
	}
	return -1 ,-1 , errors.New("doesnt exisit")

}
