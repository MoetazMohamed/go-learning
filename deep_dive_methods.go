package main 
import (
	"fmt"
)
func main7() {
	arr := []int {1, 2,3}
operationOnArr(arr, doubleNumber)
fmt.Println(arr)

operationOnArr(arr, perfectSquare)
fmt.Println(arr)

}

// passing function as parameters 
// 1. 
func doubleNumber(numb int) int {
	return 2* numb
}

func perfectSquare (a int) int {
	return a*a
}

func operationOnArr(arr []int ,fn func( int ) int ){
	for i , val := range arr {
		arr[i] = fn(val)
	}
}