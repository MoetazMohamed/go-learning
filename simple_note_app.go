package main

import (
	"bufio"
	"fmt"
	"go-learning/note"
	"go-learning/todo"
	"os"
	"strings"
)
type SaverInFile interface {
	SaveInFile() error
}
func getUserInput (prompt string ) string{
	fmt.Print(prompt)
	// CREATE  a reader that will read from command line
	reader := bufio.NewReader(os.Stdin)
	// read a tring untill a line break
	text, err := reader.ReadString('\n')

	if err != nil{
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text

}
func main3(){
	for {
	var title string = getUserInput("Please enter title of your note: ")
	var content string = getUserInput("Please enter title of your note: ")
	var todoText string = getUserData("Todo: ")

	todo , err := todo.New(todoText)
	if err != nil   {
		fmt.Printf("%v Please try again\n", err)
		fmt.Println("ckjbivebe")
		continue
	}
	note , err := note.New(title, content)
	if err != nil   {
		fmt.Printf("%v Please try again\n", err)
		
		continue
	}
	
	
	err = saveDataInfile(note)

	if err != nil {
		println(err)
		continue
	}

	err =  saveDataInfile(todo)

	if(err != nil){
		fmt.Println(err)
		continue
	}

	return
}
}

func saveDataInfile(saver SaverInFile) error{
	err := saver.SaveInFile();
	if(err != nil){
		return err
	}

	return nil
}