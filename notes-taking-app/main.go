package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"example.com/note/note"
)
func main(){
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil{
		fmt.Println(err)
		return
	}

	userNote.Display()

	err = userNote.Save()

	if err != nil{
		fmt.Println("Saving the note failed!")
		return
	}

	fmt.Println("Saving the  note succedded!")
}

func getUserInput(prompt string) (string){
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text , err :=reader.ReadString('\n')

	if err != nil{
		return ""
	}

	// now text has the \n in the txt
	// to remove it
	text = strings.TrimSuffix(text, "\n")
	// \r sometimes but not always exists
	text = strings.TrimSuffix(text, "\r")
	return text
}

func getNoteData() (string, string){
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}
