package task

import (
	"log"
	"time"
)

type Task struct {
	Name        string    `json:"name"`
	IsUrgent    bool      `json:"isUrgent"`
	IsImportant bool      `json:"isImportant"`
	Deadline    time.Time `json:"deadline"`
}

func (t *Task) IsOverdue() bool {
	return time.Now().After(t.Deadline)
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

func FilterImportant(tasks []Task) []Task {
	importantTasks := make([]Task, 0)

	for _, t := range tasks {
		if t.IsImportant {
			importantTasks = append(importantTasks, t)
		}
	}

	return importantTasks
}

func FilterUrgent(tasks []Task) []Task {
	importantTasks := make([]Task, 0)

	for _, t := range tasks {
		if t.IsUrgent {
			importantTasks = append(importantTasks, t)
		}
	}

	return importantTasks
}

func FilterNotImportant(tasks []Task) []Task {
	importantTasks := make([]Task, 0)

	for _, t := range tasks {
		if !t.IsImportant {
			importantTasks = append(importantTasks, t)
		}
	}

	return importantTasks
}

func FilterNotUrgent(tasks []Task) []Task {
	importantTasks := make([]Task, 0)

	for _, t := range tasks {
		if !t.IsUrgent {
			importantTasks = append(importantTasks, t)
		}
	}

	return importantTasks
}
