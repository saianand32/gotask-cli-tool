package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/saianand32/go_todo_cli/internal/filestorage"
	"github.com/saianand32/go_todo_cli/internal/todo"
)

func main() {
	// Initialize the file storage system
	fs, err := filestorage.New()
	if err != nil {
		fmt.Println("Error loading store:", err)
		os.Exit(1)
	}

	fmt.Println("999")
	// Initialize an empty Todos struct
	todos := &todo.Todos{}

	// Define flags for the command-line arguments
	add := flag.Bool("add", false, "add a new todo")
	ls := flag.Bool("ls", false, "list all tasks")
	flag.Parse() // Parse the flags

	// Determine which action to take based on the flags
	switch {
	case *add:
		// Check if enough arguments were provided
		if len(flag.Args()) < 2 {
			fmt.Println("Error: Please provide a group name and a task description.")
			return
		}

		group := flag.Arg(0) // Get the first positional argument (group name)
		task := flag.Arg(1)  // Get the second positional argument (task description)

		_, err := todos.Add(fs, group, task) // Pass the FileStorage instance
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Todo added successfully!")
		}

	case *ls:
		// Check if enough arguments were provided
		if len(flag.Args()) < 1 {
			fmt.Println("Error: Please provide a group name to list tasks.")
			return
		}

		group := flag.Arg(0) // Get the first positional argument (group name)

		todos.Print(fs, group) // Pass the FileStorage instance

	default:
		fmt.Println("Error: Please specify a command (--add or --ls).")
	}
}
