package task

import (
	"log"
)

type Task struct {
	Name        string `json:"name"`
	IsUrgent    bool   `json:"isUrgent"`
	IsImportant bool   `json:"isImportant"`
}

var (
	store Store = &FileStore{data: ".local/tasks.txt"}
)

func Add(t Task) {
	if err := store.Add(t); err != nil {
		log.Fatalf("Failed to add task %v with err: %s", t, err)
	}
}

func List() []Task {
	tasks, err := store.ReadAll()
	if err != nil {
		log.Fatalf("Failed to list tasks with err: %s", err)
	}
	return tasks
}
