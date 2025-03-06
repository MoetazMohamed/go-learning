package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	age       int
	createdAt time.Time
}

type Admin struct {
	email string 
	password string 
	 User
}

func (a Admin)GetAdminEmail() string{
	return a.email
}

func NewAdmin(email string, password string , user *User) (*Admin, error){
	if(email == "" || password == "" || user == nil){
		return nil, errors.New("invalid admin")
	}

	return &Admin {
	User : *user,
	email : email, 
	password: password,
	} , nil
}


func New(firstName string, lastName string, birthDate string, age int) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" || age < 1 {
		return nil, errors.New("user is invalid")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func  (u *User) GetFirstName() string{
	return u.firstName
}

func (u *User) GetLastName() string {
	return u.lastName
}

func (u *User) GetBirthDate() string{
	return u.birthDate
}

func (user *User) OutputUserDetails() {
	// if u thought about in c++ this is not allowed user.firstName
	// In C++ u need to do something like user->firstName
	// In go user.firstName
	fmt.Println((*user).firstName, user.lastName, user.birthDate)
}

func (user *User) Update(firstName string, lastName string, birthDate string, age int) {
	if age < 18 {
		fmt.Println("User u trying to create is a minor")
		fmt.Println("Update failed")
	} else {
		user.firstName = firstName
		user.lastName = lastName
		user.birthDate = birthDate
		user.age = age
	}
}
