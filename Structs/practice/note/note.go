package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Note struct {
	Title       string `json:"title"`
	NoteContent string `json:"note_content"`
}

func NewNote(title, noteContent string) *Note {
	return &Note{
		Title:       title,
		NoteContent: noteContent,
	}
}

func (note *Note) PrintContent() {
	fmt.Println(note.NoteContent)
}

func (note Note) WriteContent() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	os.WriteFile(filename, json, 0644)
	return nil
}
