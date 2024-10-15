package todo

import (
	"fmt"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/saianand32/go_todo_cli/internal/filestorage"
	"github.com/saianand32/go_todo_cli/internal/helper"
	"github.com/saianand32/go_todo_cli/internal/models"
)

type Todos []models.Item

// Add appends a new todo item to the Todos slice and writes it to the specified file.
// It takes the FileStorage instance, the group name, and the task description as parameters.
// If a todo with the same group and task already exists, it returns an error.
func (t *Todos) Add(fs *filestorage.FileStorage, task string) error {

	group, err := GetCurrentGroup(fs)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/%s.json", fs.DataFolder, group)
	data, err := fs.Read(fileName)
	if err != nil {
		return err
	}
	*t = append(*t, data...)

	id, err := helper.GenerateCryptoID()
	if err != nil {
		return err
	}

	// Check for duplicates
	for _, v := range data {
		if v.Group == group && v.Task == task {
			return fmt.Errorf("previous task with the same name already exists")
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
	err = fs.Write(fileName, *t)
	if err != nil {
		return err
	}
	return nil
}

// Complete marks a todo item as completed by setting the Done field to true
// and updating the CompletedAt field with the current time.
// It searches for the todo item using its ID and returns an error if not found.
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

// Print displays the todos in a formatted table.
// It reads the todos from the specified file and appends them to the Todos slice.
// The table includes a header for the group name and columns for UUID, task, completion status,
// and created/completed timestamps. If there are any errors during file reading, it prints the error message.
func (t *Todos) Print(fs *filestorage.FileStorage) error {

	group, err := GetCurrentGroup(fs)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/%s.json", fs.DataFolder, group)
	data, err := fs.Read(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	*t = append(*t, data...)

	table := simpletable.New()

	groupHeaderCell := &simpletable.Cell{
		Align: simpletable.AlignCenter,
		Text:  fmt.Sprintf("Group: %s", group),
	}
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 6, Text: groupHeaderCell.Text},
		},
	}

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("group: %s", strings.ToUpper(group))},
			{Align: simpletable.AlignCenter, Text: "uuid"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		task := helper.Blue(item.Task)
		done := helper.Blue("no")
		circle := "\033[33m●\033[0m" //yellow circle
		completedAt := "pending"
		if item.Done {
			task = helper.Green(item.Task)
			done = helper.Green("yes")
			circle = "\033[32m●\033[0m" //green circle
			completedAt = item.CompletedAt.Format(time.RFC822)
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: circle},
			{Text: item.Id},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 6, Text: helper.Red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
	return nil
}

// CountPending returns the number of todo items in the Todos slice that are not marked as done.
// It iterates through each todo and increments the count for those that are still pending.
func (t *Todos) CountPending() int64 {
	var count int64
	for _, todo := range *t {
		if !todo.Done {
			count++
		}
	}
	return count
}
