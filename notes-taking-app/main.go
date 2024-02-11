package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"example.com/note/note"
)

// creating interface
type saver interface{
	Save() error
}

// generic, flexible and reusable code
func savaData(data saver) error{
	err := data.Save()

	if err != nil{
		fmt.Println("Saving the data failed!")
		return err
	}

	fmt.Println("Saving the  note succedded!")
	return nil
}

// Any value allowed type interface{}
func printSpmeThing(value interface{}){
	fmt.Print(value)
}

// generic function
/*
func add [T any] (a, b T) {
	
}
*/
// t is int or float64 or string only
func add [T int | float64 | string] (a, b T){
	return a + b
}
func main(){
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil{
		fmt.Println(err)
		return
	}

	userNote.Display()
/*
this is the code before we usr interface
	err = userNote.Save()

	if err != nil{
		fmt.Println("Saving the note failed!")
		return
	}

	fmt.Println("Saving the  note succedded!")
	*/
	// it is useful if the same logic is repeated for Note and TODO for example
	err = savaData(userNote)
	if err != nil{
		return
	}
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
