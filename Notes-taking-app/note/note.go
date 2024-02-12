package note

import (
	"errors"
	"fmt"
	"time"
	"encoding/json"
	"os"
	"strings"
)

type Note struct{
	// it is nice to have diferent keys in output json file rather those in the struct
	Title string		`json:"title"` // struct tags
	Content string		`json:"content"`
	CreatedAt time.Time	`json:"created_at"`
}

func New (title, content string) (Note, error){
	if title == "" || content == ""{
		return Note{}, errors.New("invalid input.")
	}
	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display(){
	fmt.Printf("Your note title: %v has the following content:\n\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	// convert the struct to json format
	// data must e publicly available
	json , err := json.Marshal(note)

	if err != nil{
		return err
	}
	return os.WriteFile(fileName, json , 0644)
}