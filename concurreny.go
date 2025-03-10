package main 

import("fmt"
"time")

func greet(x *int) {
	fmt.Println("Greet " , *x)
	*x+=1
}
func slowGreet(){
time.Sleep(3 * time.Second)
fmt.Println("slow greet")
}

func sender(arg string , ch chan string){
	ch <- arg
}


func main(){
	  x := 10
	// why if we put all go routines nothing is executed 
	// Because actually what happens is that when you dispatch the methods the main function ends 
	// Then you go out of the main function
	// This is very important to think about 
	
	// go slowGreet()
	
	go greet(&x) // I think why we can see greet 10 in console is
	// the main thread waits for 1000 in this time the method finsihes 
	// execution before the main ends
	time.Sleep(1000)
	
	 var ch []chan string // this initializes 
	 // the channel correctly gives it memory  and set up channel 
	 // var c chan string ==> this is nil with no initilization
	ch = append(ch, make(chan string))
	 go sender("Moetaz", ch[0])
	 ch = append(ch , make(chan string))
	go sender("Moetaz", ch[1])
	ch = append(ch , make(chan string))
	go sender("Moetaz", ch[2])
	ch = append(ch , make(chan string))
	go sender("Moetaz", ch[3])

	for _, val := range ch {
		fmt.Println(<-val)
	}
	
	 //  go func (){
	// 	// var name := fetch data of user by id 
	// 	name := "moetaz"
	// 	ch <- name
	// } ()
	// data := <- ch
	// close(ch)
	// fmt.Println("data received from the channel ", data)
}