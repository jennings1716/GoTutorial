package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func NewTodo(text string) *Todo {
	return &Todo{
		Text: text,
	}
}

func (todo *Todo) PrintContent() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	os.WriteFile(filename, json, 0644)
	return nil
}
