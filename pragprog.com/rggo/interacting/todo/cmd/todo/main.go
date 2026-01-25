package main

import (
	"fmt"
	"os"
	"pragprogcom/rggo/interacting/todo"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	// Define an items list
	l := &todo.List{}

	// Use the Get method to read todo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on the number of arguments
	switch {
	// For no extra arguments, print the list
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// else, concatenate all arguments with as space an
	// add to the list as an item and save it
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
