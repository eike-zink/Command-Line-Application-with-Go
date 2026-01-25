package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// item struct respresents a todo item
type item struct {
	Task       string
	Done       bool
	CreatedAt  time.Time
	CompeledAt time.Time
}

// List represents a list of todo items
type List []item

// Add creates a new todo item and appends it to the list
func (l *List) Add(task string) {
	t := item{
		Task:       task,
		Done:       false,
		CreatedAt:  time.Now(),
		CompeledAt: time.Time{},
	}
	*l = append(*l, t)
}

// Complete method marks a todo item as completed by
// setting Done = true and CompeledAt to the current time
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exists", i)
	}

	// Adjusting index for 0 based index
	ls[i-1].Done = true
	ls[i-1].CompeledAt = time.Now()

	return nil
}

// Delete method deletes a todo item from the list
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exists.", i)
	}

	// Adjusting index for 0 based index
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save methode encodes the list as JSON and save it
// using the provided file name
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name,
// decode the JSON data and parse it into a List
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}
