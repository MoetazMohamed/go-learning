# Go Learning Notes - 2025-02-02

## What We Learned

Go is organized as follows:

* Module has many packages
* Package has many files

---------------------

* A file cannot be created without a package.
* **Format:**

    ```go
    package main

    import (
        // Import statements
        // ...
    )

    func main() {
        // Your code
    }
    ```

1.  `main` package is a special package. It tells Go that this is the entry point of your application.
2.  `main` method is very important because it is the method that executes when you run the project.
3.  **Build:** You need a `main` module to build it, and this will create an executable file that you can run on any machine.
4.  **Commands:**
    * `go mod init {name of module}`: Create module
    * `go build`: Create executable
    * `go run .`: Run executable

## Variable Declaration

* `var`, `const`
* **Format:**
    * `var variable Type = Value` (if you want to specify the type)
    * `variable := Value` (if you want the type of variable to be inferred)
* `fmt`:
    * `fmt.Print`
    * `fmt.Scan()`: You can pass references or pointers.

------------------

## How to Format Text Using the `fmt` Package

* We can use `Printf` and `Print`.
* Inside, we can use placeholders:
    * `%v`: For the value of the variable.
    * `%T`: For the type of the variable.

## Formatting Strings

* There is an `fmt` method called `Sprintf()`.
* You can format a string using this library.
* **Example:**

    ```go
    greeting := fmt.Sprintf("Hello my name %v. I appreciate that all of you came today and promise it will be exceptional", name)
    ```
### Control Structures 
For loops are everything 
for is a while, for and can be used as infinite loop
```
for :
for i:= 0 ; i< 10 ; i++ {
   fmt.printf("%.f", i) 
}

while : 
i: = 0;
for i <= 10{
    fmt.printf("%.f", i)
    i++ 
}

infinite loop 
for{
    break;
}
```
### Error Handling 
In GO, A method can return multiple things. 
Syntax
``` 
func divide(x, y int) (float64 , error){
    if(y == 0 ){
        return infinitie , errors.New("message")
    }

    return x / y, nil
}

quotient, err = divide(x,y)
if err !=  nil{
    fmt.print(quotient)
}else {
    return
}
```

### Why we use packages 
* We can organize the file int the same package. This is like just dividing you methods into different files. You dont need to import to use a method in other file 

* You can have different package. Then a package can be imported and used in different places and different projects u just need to import it. you dont import only name but the path that you should get go.mod
* Exported functions in go must be capital letter surprise mother fuckers

### Structs 
It is very similar to classes 
We use it to group related fields together and we can create method that will manipulate such data
####   How to create structs 
```
package main 
import (
    "time"
)
type User struct{
 firstName string
 lastName string 
 birthDate string
 age int 
 createdAt time.Time
}

func main(){
    // how to initialize a user using struct
    // we can use this when we dont know the order of attributes in struct
    user:=  User {
        firstName : "Moetaz",
        lastName : "Mohamed",
        birthDate : "20/02/2002,
        createdAt : time.Now(),
    }

    var user2 User =  User{
        "Moetaz",
        "Mohamed",
        "20/02/2002",
        20,
        time.Now(),

    } 

}
```

##### Methods for structs 
* In order to add a method to a struct, you need to pass this of the class refrence
* you can pass it as parameter but then you can call the method and pass the object 
* You can use a receiver. If u used this, you can call the method directly from the the object using the dot notation
```
// method to update user 
func (*User u) UpdateUser(firstName, lastName string){
    u.firstName = "krnkrn"
    u.lastName = "nfri"
}
// if u passed User only it will not update the user that called the function because it will create a new one but here you are passing a pointer.
```
#### Encapsulation
* To encapsulate attributes in a struct you can make them small letters then you cant access them out of the file 
* To make them public make the attributes' first letter capital and then they are public

#### Constructors 
* To create a contructor, u need to create a function call New. This function should return the struct type. It might also return structType, error.
* We can add validation in the constructor 

```
func New(firstName, lastName, birthDate string) User, nil{
if(firstName == "", lastName == "" , birthDate == ""){
    retrun nil , errors.New("User is invalid")
}

return User {
    firstName : firstName,
    lastName : lastName, 
    birthDate : birthDate,
}, nil
}
```

####   Expanding Structs
* The whole idea is that you can embed a struct inside a struct 
* To do this, u can just add one of the attributes as struct type make it public and you have acccess to the methods of the embeded struct 

