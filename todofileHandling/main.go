package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const FILE_Name = "todos.json"

func checkFileExists() bool {
	_, err := os.Stat(FILE_Name)
	return err == nil
}

type Todo struct {
	Index int
	Name  string
	Done  bool
}

func main() {
	fmt.Println("Welcome to Todos")
	reader := bufio.NewReader(os.Stdin)

	goOn := true
	todos, _ := loadTodos()
	for goOn {
		fmt.Println("Enter 1 for add todos, 2 for get all todos")
		ch, _ := reader.ReadString('\n')
		switch strings.TrimSpace(ch) {
		case "1":
			Name, _ := reader.ReadString('\n')
			newTodo := Todo{
				Index: len(todos) + 1,
				Name:  Name,
				Done:  false,
			}
			todos = append(todos, newTodo)
			saveTodos(todos)

		case "2":
			fmt.Println("Your todos")
			fmt.Println(todos)

		default:
			fmt.Println("Sorry wrong input")
			goOn = false
		}
	}
}

func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(FILE_Name, data, 0644)
}

func loadTodos() ([]Todo, error) {
	if _, err := os.Stat(FILE_Name); os.IsNotExist(err) {
		return []Todo{}, nil
	}

	data, err := os.ReadFile(FILE_Name)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}
