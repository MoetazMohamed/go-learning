package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)
type Todo struct {
	Text string `json:text`
}

//constructor 

func New(text string) (*Todo, error){
	fmt.Print(text)
	if(text == ""){
		fmt.Print("da5alt")
		return nil , errors.New("Invalid Todo item")
	}

	fmt.Print("md5al4")

	return &Todo{
		Text : text,
	} , nil
}

func (t *Todo ) Update(text string) error {
	if(text == ""){
		return errors.New("Update failed")
	}

	t.Text = text
	return nil

}

func (t *Todo) SaveInFile() error{
	todoJson, err := json.Marshal(t)
	fileName := "todo.json"
	if err != nil {
		return err
	}

	return  os.WriteFile(fileName, []byte(todoJson), 0644)
}