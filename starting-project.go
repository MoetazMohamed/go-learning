package main

import (
	"bufio"
	"fmt"
	"go-learning/user"
	"os"
	"strings"
)

// this how u create a struct to use in the future

// receiver parameter is exactly like parameters
// if u just used user User then whenever u update the original struct
// u dont change it because u are just changing a copy of the struct
// u need to use the pointer not a value

func main5() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	// create a variable of type user
	user1, err := user.New(firstName, lastName, birthdate, 20)
	if err != nil {
		fmt.Println(err)
		return
	}
	// var user2 User = User{
	// 	firstName,
	// 	lastName,
	// 	birthdate,
	// 	23,
	// 	time.Now(),
	// }
	user1.OutputUserDetails()
	user1.Update("Ahmed", "Yasser", "03/08/1976", 27)
	user1.OutputUserDetails()
	user1.GetFirstName()

	admin, err:= user.NewAdmin("3ezo.yasser@gmail.com", "123", user1)

	if(err != nil){
		fmt.Println(err)
	}

	fmt.Printf("ADMIN IS: %v %v", admin.GetFirstName(), admin.GetLastName())
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var userData string
	reader := bufio.NewReader(os.Stdin)
	userData, _ = reader.ReadString('\n')

	userData = strings.TrimSuffix(userData, "\n")
	userData =  strings.TrimSuffix(userData, "\r")
	return userData
}
