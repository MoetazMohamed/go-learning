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