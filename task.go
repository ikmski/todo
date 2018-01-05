package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type TaskStatus int

const (
	TaskStatusTodo TaskStatus = 0
	TaskStatusDone TaskStatus = 1
)

type Task struct {
	ID     int        `yaml:"id"`
	Title  string     `yaml:"title"`
	Detail string     `yaml:"detail"`
	Status TaskStatus `yaml:"status"`
}

type Tasks struct {
	Tasks []Task `yaml:"tasks"`
}

func newTask(file string) *Task {

	t := new(Task)
	t.ID = issueTaskID(file)
	t.Status = TaskStatusTodo

	return t
}

func loadTasks(file string) (*Tasks, error) {

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var t Tasks
	err = yaml.Unmarshal(buf, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (t *Task) save(file string) error {

	ts, err := loadTasks(file)
	if err != nil {
		return err
	}

	ts.Tasks = append(ts.Tasks, *t)

	return ts.save(file)
}

func (t *Tasks) save(file string) error {

	d, err := yaml.Marshal(t)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, d, 0644)
	if err != nil {
		return err
	}

	return nil
}

func issueTaskID(file string) int {

	ts, err := loadTasks(file)
	if err != nil {
		return 0
	}

	id := 0
	for _, t := range ts.Tasks {
		if t.ID > id {
			id = t.ID
		}
	}
	id++

	return id
}

func (t *Task) interactiveEdit() error {

	fmt.Printf("title:[%s] ", t.Title)
	fmt.Scanln(&t.Title)

	fmt.Printf("Detail:[%s] ", t.Detail)
	fmt.Scanln(&t.Detail)

	return nil
}

func (ts *Tasks) interactiveDelete(id int) (bool, error) {

	tasks := ts.Tasks

	idx := -1
	for i, t := range ts.Tasks {
		if t.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false, fmt.Errorf("task not found")
	}

	s := ""
	fmt.Printf("Delete task %03d: %s ? (y/N) ", tasks[idx].ID, tasks[idx].Title)
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {

		result := []Task{}
		for _, t := range ts.Tasks {
			if t.ID != id {
				result = append(result, t)
			}
		}

		ts.Tasks = result
		return true, nil
	}

	return false, nil
}

func (ts *Tasks) interactiveDone(id int) (bool, error) {

	tasks := ts.Tasks

	idx := -1
	for i, t := range ts.Tasks {
		if t.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false, fmt.Errorf("task not found")
	}

	if tasks[idx].Status == TaskStatusDone {
		return false, fmt.Errorf("task status is already DONE")
	}

	s := ""
	fmt.Printf("Done task %03d: %s ? (y/N) ", tasks[idx].ID, tasks[idx].Title)
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {

		tasks[idx].Status = TaskStatusDone
		return true, nil
	}

	return false, nil
}

func (ts *Tasks) interactiveUndone(id int) (bool, error) {

	tasks := ts.Tasks

	idx := -1
	for i, t := range ts.Tasks {
		if t.ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false, fmt.Errorf("task not found")
	}

	if tasks[idx].Status == TaskStatusTodo {
		return false, fmt.Errorf("task status is already TODO")
	}

	s := ""
	fmt.Printf("Undone task %03d: %s ? (y/N) ", tasks[idx].ID, tasks[idx].Title)
	fmt.Scanln(&s)
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {

		tasks[idx].Status = TaskStatusTodo
		return true, nil
	}

	return false, nil
}
