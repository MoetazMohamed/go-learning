package main

import (
	"fmt"
	"go-learning/note"
	"bufio"
	"os"
	"strings"
)

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
func main(){
	for {
	var title string = getUserInput("Please enter title of your note: ")
	var content string = getUserInput("Please enter title of your note: ")


	note , err := note.New(title, content)

	if err != nil {
		fmt.Println("Please try again")
		continue
	}

	err = note.SaveNoteInFile()

	if err != nil {
		println(err)
		continue
	}

	return
}
}