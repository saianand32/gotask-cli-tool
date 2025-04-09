# Go Todo CLI

A simple, powerful command-line todo list application built with Go that helps you organize tasks in groups.

## Features

- Add todo items with descriptions
- Mark tasks as complete
- List all your tasks
- Organize todos in separate groups
- Manage groups (create, list, delete, clear)
- Persistent storage of your tasks

## Installation

```bash
# Clone the repository
git clone https://github.com/saianand32/gotask-cli-tool.git

# Navigate to the project directory
cd gotask-cli-tool

# Build the application
go build -o gotask
```



## Commands Details

| Command | Description | Example |
|---------|-------------|---------|
| `add` | Add a new todo task | `./todo add "Complete project report"` |
| `done` | Mark a todo as complete | `./todo done 2` |
| `ls` | List all todos in the current group | `./todo ls` |
| `usegrp` | Create or switch to a group | `./todo usegrp personal` |
| `showgrp` | Display all available groups | `./todo showgrp` |
| `dropgrp` | Delete a group and its todos | `./todo dropgrp old_project` |
| `truncategrp` | Remove all todos from a group | `./todo truncategrp work` |



## ProjectStructure
```bash
gotask-cli-tool/
├── internal/
│   ├── constants/
│   │   └── constants.go          # List of valid commands and other constants
│   ├── filestorage/
│   │   └── filestorage.go        # File-based storage implementation
│   └── todo/
│       └── todo.go               # Core logic for managing todos and groups
├── main.go                       # CLI entry point
└── README.md 
```