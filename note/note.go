package note

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Create_at time.Time `json:"created_At"`
}

func New(title, content string) (*Note, error) {
	if title == "" {
		return nil, errors.New("Note doesnt have a title")
	}

	return &Note{
		Title:     title,
		Content:   content,
		Create_at: time.Now(),
	}, nil
}

func (n *Note) UpdateNote(title, content string) error {
	if title == "" {
		return errors.New("update failed title is missing")
	}
	n.Title = title
	n.Content = content
	return nil
}

func (n *Note) GetTitle() string {
	return n.Title
}

func (n *Note) GetContent() string {
	return n.Content
}

func (n Note) SaveInFile() error {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName =  strings.ToLower(fileName) + ".json"
	jsonNote, err := json.Marshal(n)

	if err != nil {
		return err
	}
	
	return os.WriteFile(fileName, jsonNote, 0644)
}