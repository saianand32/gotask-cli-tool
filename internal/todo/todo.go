package todo

import (
	"fmt"
	"time"

	"github.com/saianand32/go_todo_cli/internal/filestorage"
	"github.com/saianand32/go_todo_cli/internal/helper"
	"github.com/saianand32/go_todo_cli/internal/models"
)

type Todos []models.Item

func (t *Todos) Add(fs *filestorage.FileStorage, group, task string) (bool, error) {

	fileName := fmt.Sprintf("%s/%s.json", fs.DataFolder, group)
	data, err := fs.Read(fileName)
	if err != nil {
		return false, err
	}
	*t = append(*t, data...)

	id, err := helper.GenerateCryptoID()
	if err != nil {
		return false, err
	}

	// Check for duplicates
	for _, v := range data {
		if v.Group == group && v.Task == task {
			return false, fmt.Errorf("previous task with the same name already exists")
		}
	}

	todo := models.Item{
		Id:          id,
		Group:       group,
		Task:        task,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
	isSuccess, err := fs.Write(fileName, *t)
	if err != nil {
		return false, err
	}
	return isSuccess, nil
}

func (t *Todos) Complete(fs *filestorage.FileStorage, id string) (bool, error) {
	for i, todo := range *t {
		if todo.Id == id {
			(*t)[i].Done = true
			(*t)[i].CompletedAt = time.Now()
			return true, nil
		}
	}
	return false, fmt.Errorf("todo with id %s not found", id)
}

func (t *Todos) Delete(fs *filestorage.FileStorage, id string) (bool, error) {
	index := -1
	for i, todo := range *t {
		if todo.Id == id {
			index = i
			break
		}
	}
	if index == -1 {
		return false, fmt.Errorf("todo with id %s not found", id)
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
	return true, nil
}
