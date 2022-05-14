package task

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Store interface {
	Add(t Task) error
	ReadAll() ([]Task, error)
}

type FileStore struct {
	data string
}

func (fs *FileStore) Add(t Task) error {
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("Failed to marshal the task %v with error %s", t, err)
		return err
	}
	bytes = append(bytes, byte('\n'))

	f, err := os.OpenFile(fs.data, os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		log.Printf("Failed to open the file for tasks with err: %s", err)
		return err
	}
	defer f.Close()

	if _, err := f.Write(bytes); err != nil {
		log.Printf("Failed to persist the task %v with error %s", t, err)
		return err
	}

	return nil
}

func (fs *FileStore) ReadAll() ([]Task, error) {
	f, err := os.Open(fs.data)
	if err != nil {
		log.Printf("Failed to open the file for tasks with err: %s", err)
		return make([]Task, 0), err
	}
	defer f.Close()

	tasks := make([]Task, 0)
	// Scanner's capacity is capped to lines under 64k characters in length. Read this for more information:
	// https://stackoverflow.com/a/16615559
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var task Task
		bytes := scanner.Bytes()
		err := json.Unmarshal(bytes, &task)
		if err != nil {
			log.Printf("Skipping task %s because it cannot be unmarshalled with error: %s", bytes, err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
