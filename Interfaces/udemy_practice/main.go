package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practiceinterfaces/note/note"
	"example.com/practiceinterfaces/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	PrintContent()
}

type outputtable interface {
	saver
	displayer
}

func main() {

	var title string = getNoteData("Enter the title of Note")
	var content string = getNoteData("Enter the content of Note")

	var text string = getNoteData("Enter the text of Todo")

	var testNote = note.NewNote(title, content)
	var testTodo = todo.NewTodo(text)
	testNote.PrintContent()

	saveAndDisplay(testNote)
	saveAndDisplay(testTodo)

	saveData(testNote)
	saveData(testTodo)

	displayData(testNote)
	displayData(testTodo)

	printSomething(1)
	printSomething(1.2)
	printSomething("TEST")
}

func saveData(data saver) {
	data.Save()
}

func displayData(data displayer) {
	data.PrintContent()
}

func saveAndDisplay(data outputtable) {
	data.Save()
	data.PrintContent()
}

func getNoteData(promptText string) string {
	fmt.Print(promptText)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading")
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	return text
}

func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("float")
	}
}
