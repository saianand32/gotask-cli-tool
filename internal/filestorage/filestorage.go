package filestorage

import (
	"encoding/json"
	"os"

	"github.com/saianand32/go_todo_cli/internal/constants"
	"github.com/saianand32/go_todo_cli/internal/helper"
	"github.com/saianand32/go_todo_cli/internal/models"
)

type FileStorage struct {
	BasePath     string
	DataFolder   string
	ConfigFolder string
	GroupFile    string
}

func New() (*FileStorage, error) {
	folders := constants.Folders
	files := constants.Files

	for _, folder := range helper.GetMapValues(folders) {
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			// Create the directory if it does not exist
			err := os.MkdirAll(folder, 0777)
			if err != nil {
				return nil, err
			}
		}
	}
	for _, file := range helper.GetMapValues(files) {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			// Create the file if it does not exist
			f, err := os.Create(file)
			if err != nil {
				return nil, err
			}
			defer f.Close() // Ensure the file is closed after creation
		}
	}
	return &FileStorage{
		BasePath:     folders["StoreFolder"],
		ConfigFolder: folders["ConfigFolder"],
		DataFolder:   folders["DataFolder"],
		GroupFile:    files["GroupFile"],
	}, nil
}

func (fs *FileStorage) Read(fileName string) ([]models.Item, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []models.Item{}, nil
	}
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	var items []models.Item
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&items) // Unmarshal JSON into the slice
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (fs *FileStorage) Write(fileName string, data []models.Item) (bool, error) {
	file, err := os.Create(fileName) // os.Create will either create or truncate the file
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Encode the new data as JSON and write it to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data) // Marshal the data to JSON and write it
	if err != nil {
		return false, err
	}

	return true, nil
}
