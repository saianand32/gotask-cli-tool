package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/saianand32/go_todo_cli/internal/constants"
	"github.com/saianand32/go_todo_cli/internal/filestorage"
	"github.com/saianand32/go_todo_cli/internal/todo"
)

func main() {
	// Initialize the file storage system
	fs, err := filestorage.New()
	if err != nil {
		fmt.Println("Error loading store:", err)
		return
	}

	// Initialize an empty Todos struct
	todos := &todo.Todos{}

	// Check if there are enough arguments
	if len(os.Args) < 2 {
		fmt.Println("error: Please specify a command")
		return
	}

	// Get the command (first positional argument)
	command := os.Args[1]
	if !slices.Contains(constants.ValidCommands, command) {
		fmt.Println("error: invalid command", command)
		return
	}

	// Handle commands
	switch command {
	case "add":
		// Ensure we have enough arguments for the add command
		if len(os.Args) < 2 {
			fmt.Println("error: Please provide a group name and a task description.")
			return
		}
		task := os.Args[2] // Get the third positional argument (task description)

		err := todos.Add(fs, task) // Pass the FileStorage instance
		if err != nil {
			fmt.Println("error:", err)
			return
		} else {
			fmt.Println("Todo added successfully!")
		}

	case "ls":
		// Ensure we have enough arguments for the ls command
		if len(os.Args) < 1 {
			fmt.Println("Error: Please provide a group name to list tasks.")
			return
		}

		err := todos.Print(fs) // Pass the FileStorage instance
		if err != nil {
			fmt.Println(err)
			return
		}
	case "usegrp":
		// Ensure we have enough arguments for the ls command
		if len(os.Args) < 2 {
			fmt.Println("Error: Please provide group name to use/create")
			return
		}
		group := os.Args[2]
		err := todo.CreateGroup(fs, group) // Pass the FileStorage instance
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Using: %s\n", group)

	case "showgrp":
		err := todo.ListGroups(fs) // Pass the FileStorage instance
		if err != nil {
			fmt.Println(err)
			return
		}
	case "dropgrp":
		group := os.Args[2]
		err := todo.DropGroup(fs, group)
		if err != nil {
			fmt.Println(err)
			return
		}

	default:
		fmt.Println("Error: Unknown command.")
	}

}
