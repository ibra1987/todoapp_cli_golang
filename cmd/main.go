package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	Title     string
	Completed bool
}

var todos []Todo

func main() {
	var title string
	todo := &Todo{}
	todoList := todo.GetTodos()
	quit := false

	scanner := bufio.NewScanner(os.Stdin)
	var choiceMap = map[string]string{
		"1": "Add a new todo",
		"2": "Print the todo list",
		"3": "Delete a todo",
		"4": "Quit",
	}

	fmt.Println("----------------------------------------------")
	fmt.Println("Todo List App")
	fmt.Println("----------------------------------------------")

	for !quit {

		PrintMenu(choiceMap)

		if scanner.Scan() {
			choice := scanner.Text()
			switch choice {
			case "1":
				fmt.Println("PLease enter the title of the todo:")
				if scanner.Scan() {
					title = scanner.Text()
					todo.NewTodo(title)
				}

			case "2":
				PrintTodos(todoList)

			case "3":
				fmt.Println("Please enter the id of the todo you want to delete")
				if scanner.Scan() {
					id, err := strconv.Atoi(scanner.Text())
					if err != nil {
						fmt.Println(err)

					}
					if id < 1 || id > len(*todoList) {
						fmt.Println("Invalid Id")

					}
					todo.deleteTodo(todoList, id)
				}

			case "4":
				quit = true
			default:
				fmt.Println("Invalid input")

			}
		} else if scanner.Err() != nil {
			fmt.Println(errors.New("error reading input"))
		}

	}

}

func (todo *Todo) NewTodo(title string) (*[]Todo, error) {

	if title == "" {
		return nil, errors.New("title is required")
	}

	todos = append(todos, Todo{
		Title:     title,
		Completed: false,
	})
	fmt.Println("Todo created successfully!")
	return &todos, nil

}

func (todo *Todo) GetTodos() *[]Todo {
	return &todos
}

func PrintTodos(todoList *[]Todo) {
	println("*******************************************************")
	for i, t := range *todoList {
		fmt.Printf("%d. %s : %s\n", i+1, t.Title, formatCompleted(t.Completed))
	}
	println("*******************************************************")
}

func PrintMenu(choiceMap map[string]string) {
	println()
	for key, value := range choiceMap {
		fmt.Println(key + ":" + value)
	}
	println()
}

func (todo *Todo) deleteTodo(todos *[]Todo, id int) {

	*todos = append((*todos)[:id-1], (*todos)[id:]...)
	fmt.Println("Todo deleted successfully!")
}

func formatCompleted(isCompleted bool) string {
	if isCompleted {
		return "Completed"
	} else {
		return "Not Completed"
	}
}
