package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practicestructs/note"
)

func main() {

	var title string = getNoteData("Enter the title of Note")
	var content string = getNoteData("Enter the content of Note")

	var testNote = note.NewNote(title, content)
	testNote.PrintContent()
	testNote.WriteContent()

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
