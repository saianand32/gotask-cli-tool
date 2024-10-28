package todo

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/saianand32/go_todo_cli/internal/filestorage"
	"github.com/saianand32/go_todo_cli/internal/helper"
)

// CreateGroup creates a new group by storing the group name in the GroupFile and
// creating an empty JSON file for tasks within the DataFolder.
func CreateGroup(fs *filestorage.FileStorage, group string) error {
	// Write the group name to the GroupFile.
	groupName := strings.ToLower(group)
	data := []byte(groupName)
	err := os.WriteFile(fs.GroupFile, data, 0644)
	if err != nil {
		return fmt.Errorf("couldn't write to file: %v", err)
	}

	// Create the groupName file path.
	fileName := fmt.Sprintf("%s/%s.json", fs.DataFolder, groupName)

	// Check if the file already exists.
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// If the file doesn't exist, create it with an empty JSON array.
		err := os.WriteFile(fileName, []byte("[]"), 0644)
		if err != nil {
			return fmt.Errorf("couldn't create group JSON file: %v", err)
		}
	} else if err != nil {
		// If there was another error (other than "file does not exist"), return it.
		return fmt.Errorf("error checking group file: %v", err)
	}

	// Return success if no errors occurred.
	return nil
}

func GetCurrentGroup(fs *filestorage.FileStorage) (string, error) {
	file, err := os.Open(fs.GroupFile) // Open the file
	if err != nil {
		return "", fmt.Errorf("couldn't open groups file: %v", err)
	}
	defer file.Close() // Ensure the file is closed when the function ends

	groupData, err := io.ReadAll(file) // Read the file's contents
	if err != nil {
		return "", fmt.Errorf("couldn't read groups file: %v", err)
	}

	group := strings.TrimSpace(string(groupData)) // Convert the byte slice to a string
	if group == "" {
		return "", fmt.Errorf("no group selected. use 'usegrp <group_name>' to select a group")
	}
	return group, nil
}

// ListGroups lists all available groups by scanning the DataFolder for JSON files.
// The current group is highlighted in blue when listed.
func ListGroups(fs *filestorage.FileStorage) error {
	dirEntries, err := os.ReadDir(fs.DataFolder)
	if err != nil {
		return fmt.Errorf("couldn't list groups: %v", err)
	}

	currentGroup, _ := GetCurrentGroup(fs)

	fmt.Println("Available groups:")
	for _, entry := range dirEntries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			groupName := strings.TrimSuffix(entry.Name(), ".json")

			if groupName == currentGroup {
				fmt.Println("- " + helper.Green(groupName))
			} else {
				fmt.Println("- " + groupName)
			}
		}
	}

	return nil
}

func DropGroup(fs *filestorage.FileStorage, group string) error {

	fileName := fmt.Sprintf("%s/%s.json", fs.DataFolder, strings.ToLower(group))
	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return err
	}

	cur_group, err := GetCurrentGroup(fs)
	if err != nil {
		return fmt.Errorf("fetching current group")
	}

	if strings.EqualFold(cur_group, group) {
		data := []byte("")
		err = os.WriteFile(fs.GroupFile, data, 0644)
		if err != nil {
			return fmt.Errorf("couldn't write to file: %v", err)
		}
	}
	return nil
}
